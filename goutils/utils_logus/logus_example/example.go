package logus_example

import (
	"os"
	"strings"

	"github.com/darklab8/darklab_goutils/goutils/utils_logus"
)

var (
	Slogger *utils_logus.Logger
)

func init() {
	Slogger = utils_logus.NewLogger(utils_logus.LEVEL_DEBUG, strings.ToLower(os.Getenv("GOUTILS_LOG_JSON")) == "true", false)
}
