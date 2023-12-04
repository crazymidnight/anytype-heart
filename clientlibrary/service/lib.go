package service

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/anyproto/anytype-heart/core"
	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/metrics"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/util/vcs"
)

var log = logging.Logger("anytype-mw")

var mw = core.New()

func init() {
	fixTZ()
	fmt.Printf("mw lib: %s\n", vcs.GetVCSInfo().Description())


	PanicHandler = mw.OnPanic
	metrics.Service.InitWithKeys(metrics.DefaultAmplitudeKey, metrics.DefaultInHouseKey)
	registerClientCommandsHandler(
		&ClientCommandsHandlerProxy{
			client: mw,
			interceptor: func(ctx context.Context, methodName string, actualCall func(ctx context.Context, req any) any) any {
				start := time.Now().UnixMilli()
				resp := actualCall(ctx, methodName)
				delta := time.Now().UnixMilli() - start
				metrics.SendMethodEvent(methodName, nil, resp, delta)
				return resp
			},
		})
	if debug, ok := os.LookupEnv("ANYPROF"); ok && debug != "" {
		go func() {
			http.ListenAndServe(debug, nil)
		}()
	}
}

func SetEventHandler(eh func(event *pb.Event)) {
	mw.SetEventSender(event.NewCallbackSender(eh))
}

func SetEnv(key, value string) {
	os.Setenv(key, value)
}

func SetEventHandlerMobile(eh MessageHandler) {
	SetEventHandler(func(event *pb.Event) {
		b, err := proto.Marshal(event)
		if err != nil {
			log.Errorf("eventHandler failed to marshal error: %s", err.Error())
		}
		eh.Handle(b)
	})
}