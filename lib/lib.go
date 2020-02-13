package lib

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/anytypeio/go-anytype-middleware/core"
	"github.com/anytypeio/go-anytype-middleware/pb"

	"github.com/gogo/protobuf/proto"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("anytype-mw")

var mw = core.NewMiddleware()

func init() {
	registerClientCommandsHandler(mw)
	if debug, ok := os.LookupEnv("ANYPROF"); ok && debug != "0" {
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
	}
}

func SetEventHandler(eh func(event *pb.Event)) {
	mw.SendEvent = eh
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
