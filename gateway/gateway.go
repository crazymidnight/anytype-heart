package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anytypeio/go-anytype-library/core"
	"github.com/anytypeio/go-anytype-library/logging"
)

const defaultGatewayAddr = "127.0.0.1:47800"

var log = logging.Logger("anytype-gateway")

// Host is the instance used by the daemon
var Host *Gateway

// Gateway is a HTTP API for getting files and links from IPFS
type Gateway struct {
	Node   core.Service
	server *http.Server
}

func GatewayAddr() string {
	if addr := os.Getenv("ANYTYPE_GATEWAY_ADDR"); addr != "" {
		return addr
	}

	return defaultGatewayAddr
}

// Start creates a gateway server
func (g *Gateway) Start(addr string) error {
	if g.server != nil {
		return fmt.Errorf("gateway already started")
	}

	handler := http.NewServeMux()
	g.server = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	handler.HandleFunc("/file/", g.fileHandler)
	handler.HandleFunc("/image/", g.imageHandler)

	// check port first
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// todo: choose next available port
		return err
	}

	err = listener.Close()
	if err != nil {
		return err
	}

	errc := make(chan error)
	go func() {
		errc <- g.server.ListenAndServe()
		close(errc)
	}()
	go func() {
		for {
			select {
			case err, ok := <-errc:
				if err != nil && err != http.ErrServerClosed {
					log.Errorf("gateway error: %s", err)
				}

				if !ok {
					log.Info("gateway was shutdown")
					return
				}
			}
		}
	}()

	log.Infof("gateway listening at %s", g.server.Addr)
	return nil
}

// Stop stops the gateway
func (g *Gateway) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return g.server.Shutdown(ctx)
}

// Addr returns the gateway's address
func (g *Gateway) Addr() string {
	return g.server.Addr
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}

// fileHandler gets file meta from the DB, gets the corresponding data from the IPFS and decrypts it
func (g *Gateway) fileHandler(w http.ResponseWriter, r *http.Request) {
	fileHash := r.URL.Path[len("/file/"):]
	enableCors(w)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	file, err := g.Node.FileByHash(ctx, fileHash)
	if err != nil {
		if strings.Contains(err.Error(), "file not found") {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}

	reader, err := file.Reader()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	meta := file.Meta()
	w.Header().Set("Content-Type", meta.Media)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", meta.Name))

	// todo: inside textile it still requires the file to be fully downloaded and decrypted(consuming 2xSize in ram) to provide the ReadSeeker interface
	// 	need to find a way to use ReadSeeker all the way from downloading files from IPFS to writing the decrypted chunk to the HTTP
	http.ServeContent(w, r, meta.Name, meta.Added, reader)
}

// fileHandler gets file meta from the DB, gets the corresponding data from the IPFS and decrypts it
func (g *Gateway) imageHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	imageHash := urlParts[2]
	query := r.URL.Query()

	enableCors(w)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	image, err := g.Node.ImageByHash(ctx, imageHash)
	if err != nil {
		if strings.Contains(err.Error(), "file not found") {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}
	var file core.File
	wantWidthStr := query.Get("width")
	if wantWidthStr == "" {
		file, err = image.GetFileForLargestWidth(ctx)
	} else {
		wantWidth, err2 := strconv.Atoi(wantWidthStr)
		if err2 != nil {
			http.Error(w, err2.Error(), 400)
			return
		}

		file, err = image.GetFileForWidth(ctx, wantWidth)
	}

	if err != nil {
		if strings.Contains(err.Error(), "file not found") {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}

	reader, err := file.Reader()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	meta := file.Meta()
	w.Header().Set("Content-Type", meta.Media)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", meta.Name))

	// todo: inside textile it still requires the file to be fully downloaded and decrypted(consuming 2xSize in ram) to provide the ReadSeeker interface
	// 	need to find a way to use ReadSeeker all the way from downloading files from IPFS to writing the decrypted chunk to the HTTP
	http.ServeContent(w, r, meta.Name, meta.Added, reader)
}
