package utils_logus

import (
	"os"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/logus/logus_types"
)

var (
	Log *logus.Logger
)

func init() {
	Log = logus.NewLogger(
		logus_types.LogLevel(os.Getenv("UTILS_LOG_LEVEL")),
		logus_types.EnableJsonFormat(os.Getenv("UTILS_LOG_JSON") == "true"),
		logus_types.EnableFileShowing(os.Getenv("UTILS_LOG_FILE_SHOWING") == "true"),
	)
}
