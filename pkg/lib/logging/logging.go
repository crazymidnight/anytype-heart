package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/anytypeio/any-sync/app/logger"
	"go.uber.org/zap"

	"github.com/anytypeio/go-anytype-middleware/util/vcs"
)

var defaultCfg = logger.Config{
	Production:   false,
	DefaultLevel: "WARN",
	Format:       logger.JSONOutput,
}

type LWrapper struct {
	*zap.SugaredLogger
}

func (l LWrapper) Warningf(template string, args ...interface{}) {
	l.Warnf(template, args...)
}

func Logger(system string) *zap.SugaredLogger {
	return logger.NewNamedSugared(system)
}

func LoggerNotSugared(system string) *zap.Logger {
	lg := logger.NewNamed(system)

	return lg.Logger
}

func LevelsFromStr(s string) map[string]string {
	levels := make(map[string]string)
	for _, kv := range strings.Split(s, ";") {
		strings.TrimSpace(kv)
		parts := strings.Split(kv, "=")
		var key, value string
		if len(parts) == 1 {
			key = "*"
			value = parts[0]
			levels["*"] = parts[0]
		} else if len(parts) == 2 {
			key = parts[0]
			value = parts[1]
		}
		_, err := zap.ParseAtomicLevel(value)
		if err != nil {
			fmt.Printf("Can't parse log level %s: %s\n", parts[0], err.Error())
			continue
		}
		levels[key] = value
	}
	return levels
}

func init() {
	cfg := defaultCfg
	info := vcs.GetVCSInfo()

	SetVersion(info.Version())

	if os.Getenv("ANYTYPE_LOG_NOGELF") == "1" {
		cfg.Format = logger.ColorizedOutput
	} else {
		registerGelfSink(&cfg)
	}
	cfg.Levels = logger.LevelsFromStr(os.Getenv("ANYTYPE_LOG_LEVEL"))
	cfg.ApplyGlobal()
}
