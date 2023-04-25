//go:build !nogrpcserver && !_test
// +build !nogrpcserver,!_test

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	//nolint: gosec
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/anytypeio/any-sync/app"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	"github.com/anytypeio/go-anytype-middleware/core"
	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/anytypeio/go-anytype-middleware/metrics"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pb/service"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
)

const defaultAddr = "127.0.0.1:31007"
const defaultWebAddr = "127.0.0.1:31008"
const defaultUnaryWarningAfter = time.Second * 3

// do not change this, js client relies on this msg to ensure that server is up
const grpcWebStartedMessagePrefix = "gRPC Web proxy started at: "

var log = logging.Logger("anytype-grpc-server")

func main() {
	var addr string
	var webaddr string

	fmt.Printf("mw grpc: %s\n", app.VersionDescription())
	if len(os.Args) > 1 {
		addr = os.Args[1]
		if len(os.Args) > 2 {
			webaddr = os.Args[2]
		}
	}

	if addr == "" {
		if env := os.Getenv("ANYTYPE_GRPC_ADDR"); env != "" {
			addr = env
		} else {
			addr = defaultAddr
		}
	}

	if webaddr == "" {
		if env := os.Getenv("ANYTYPE_GRPCWEB_ADDR"); env != "" {
			webaddr = env
		} else {
			webaddr = defaultWebAddr
		}
	}

	if debug, ok := os.LookupEnv("ANYPROF"); ok && debug != "" {
		go func() {
			http.ListenAndServe(debug, nil)
		}()
	}
	metrics.SharedClient.InitWithKey(metrics.DefaultAmplitudeKey)

	var stopChan = make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var mw = core.New()
	mw.EventSender = event.NewGrpcSender()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	addr = lis.Addr().String()

	webLis, err := net.Listen("tcp", webaddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	webaddr = webLis.Addr().String()
	var (
		unaryInterceptors  []grpc.UnaryServerInterceptor
		streamInterceptors []grpc.StreamServerInterceptor
	)

	if metrics.Enabled {
		unaryInterceptors = append(unaryInterceptors, grpc_prometheus.UnaryServerInterceptor)
	}
	unaryInterceptors = append(unaryInterceptors, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = mw.Authorize(ctx, req, info, handler)
		if err != nil {
			log.Errorf("authorize: %s", err)
		}
		return
	})

	// todo: we may want to change it to the opposite check with a public release
	if os.Getenv("ANYTYPE_GRPC_NO_DEBUG_TIMEOUT") != "1" {
		unaryInterceptors = append(unaryInterceptors, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			doneCh := make(chan struct{})
			start := time.Now()
			go func() {
				select {
				case <-doneCh:
				case <-time.After(defaultUnaryWarningAfter):
					trace := base64.RawStdEncoding.EncodeToString(stackAllGoroutines())
					log.With("method", info.FullMethod).With("in_progress", true).With("goroutines", trace).With("total", defaultUnaryWarningAfter.Milliseconds()).Warnf("grpc unary request is taking too long")
				}
			}()
			resp, err = handler(ctx, req)
			close(doneCh)
			if time.Since(start) > defaultUnaryWarningAfter {
				log.With("method", info.FullMethod).With("in_progress", false).With("total", time.Since(start).Milliseconds()).Warnf("grpc unary request took too long")
			}
			return
		})
	}

	grpcDebug, _ := strconv.Atoi(os.Getenv("ANYTYPE_GRPC_LOG"))
	if grpcDebug > 0 {
		decider := func(_ context.Context, _ string, _ interface{}) bool {
			return true
		}

		grpcLogger := logging.LoggerNotSugared("grpc")

		unaryInterceptors = append(unaryInterceptors, grpc_zap.UnaryServerInterceptor(grpcLogger))
		streamInterceptors = append(streamInterceptors, grpc_zap.StreamServerInterceptor(grpcLogger))
		if grpcDebug > 1 {
			unaryInterceptors = append(unaryInterceptors, grpc_zap.PayloadUnaryServerInterceptor(grpcLogger, decider))
		}
		if grpcDebug > 2 {
			streamInterceptors = append(streamInterceptors, grpc_zap.PayloadStreamServerInterceptor(grpcLogger, decider))
		}
	}

	grpcTrace, _ := strconv.Atoi(os.Getenv("ANYTYPE_GRPC_TRACE"))
	if grpcTrace > 0 {
		jLogger := jaeger.StdLogger

		cfg, err := jaegercfg.FromEnv()
		if err != nil {
			log.Fatal(err.Error())
		}
		if cfg.ServiceName == "" {
			cfg.ServiceName = "mw"
		}
		// Initialize tracer with a logger and a metrics factory
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jLogger))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer closer.Close()

		var (
			unaryOptions  []otgrpc.Option
			streamOptions []otgrpc.Option
		)

		// Set the singleton opentracing.Tracer with the Jaeger tracer.
		opentracing.SetGlobalTracer(tracer)
		if grpcTrace > 1 {
			unaryOptions = append(unaryOptions, otgrpc.LogPayloads())
		}
		if grpcTrace > 2 {
			streamOptions = append(streamOptions, otgrpc.LogPayloads())
		}

		unaryInterceptors = append(unaryInterceptors, otgrpc.OpenTracingServerInterceptor(tracer, unaryOptions...))
		streamInterceptors = append(streamInterceptors, otgrpc.OpenTracingStreamServerInterceptor(tracer, streamOptions...))
	}

	unaryInterceptors = append(unaryInterceptors, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				mw.OnPanic(r)
				resp = &pb.RpcGenericErrorResponse{
					Error: &pb.RpcGenericErrorResponseError{
						Code:        pb.RpcGenericErrorResponseError_UNKNOWN_ERROR,
						Description: "panic recovered",
					},
				}
			}
		}()

		resp, err = handler(ctx, req)
		return resp, err
	})

	server := grpc.NewServer(grpc.MaxRecvMsgSize(20*1024*1024),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
	)

	service.RegisterClientCommandsServer(server, mw)
	if metrics.Enabled {
		grpc_prometheus.EnableHandlingTimeHistogram()
		//grpc_prometheus.Register(server)
	}

	webrpc := grpcweb.WrapServer(
		server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
			return true
		}))

	proxy := &http.Server{
		Addr: webaddr,
	}

	proxy.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if webrpc.IsGrpcWebRequest(r) ||
			webrpc.IsAcceptableGrpcCorsRequest(r) ||
			webrpc.IsGrpcWebSocketRequest(r) {
			webrpc.ServeHTTP(w, r)
		}
	})

	go func() {
		server.Serve(lis)
	}()
	fmt.Println("gRPC server started at: " + addr)

	go func() {
		if err := proxy.Serve(webLis); err != nil && err != http.ErrServerClosed {
			log.Fatalf("proxy error: %v", err)
		}
	}()

	// do not change this, js client relies on this msg to ensure that server is up and parse address
	fmt.Println(grpcWebStartedMessagePrefix + webaddr)

	select {
	case <-stopChan:
		server.Stop()
		proxy.Close()
		mw.AppShutdown(context.Background(), &pb.RpcAppShutdownRequest{})
		return
	}
}

func stackAllGoroutines() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
