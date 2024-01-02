package logus_example

import (
	"os"
	"strings"

	"github.com/darklab8/darklab_goutils/goutils/logus"
)

var (
	Slogger *logus.Logger
)

func init() {
	Slogger = logus.NewLogger(logus.LEVEL_DEBUG, strings.ToLower(os.Getenv("GOUTILS_LOG_JSON")) == "true", false)
}
