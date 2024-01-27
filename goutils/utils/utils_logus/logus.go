package utils_logus

import (
	"fmt"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

var (
	Log *logus_core.Logger
)

func init() {
	Log = logus_core.NewLogger(
		logus_core.WithLogLevelStr(os.Getenv("UTILS_LOG_LEVEL")),
	)
}

func Regex(value utils_types.RegExp) logus_core.SlogParam {
	return func(c *logus_core.SlogGroup) {
		c.Params["regexp"] = fmt.Sprintf("%v", value)
	}
}
