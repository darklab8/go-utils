package utils_logus

import (
	"fmt"
	"log/slog"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/darklab8/logusgo/logcore"
)

var Log *logcore.Logger = logcore.NewLogger("goutils")

func Regex(value utils_types.RegExp) logcore.SlogParam {
	return func(c *logcore.SlogGroup) {
		c.Append(slog.String("regexp", fmt.Sprintf("%v", value)))
	}
}

func FilePath(value utils_types.FilePath) logcore.SlogParam {
	return func(c *logcore.SlogGroup) {
		c.Append(slog.String("filepath", fmt.Sprintf("%v", value)))
	}
}

func Filepaths(values []utils_types.FilePath) logcore.SlogParam {
	return logcore.Items[utils_types.FilePath](values, "filepaths")
}
