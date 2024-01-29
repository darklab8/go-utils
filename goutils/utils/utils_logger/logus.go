package utils_logger

import (
	"fmt"
	"log/slog"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/darklab8/go-typelog/typelog"
)

var Log *typelog.Logger = typelog.NewLogger("goutils")

func Regex(value utils_types.RegExp) typelog.LogType {
	return func(c *typelog.LogAtrs) {
		c.Append(slog.String("regexp", fmt.Sprintf("%v", value)))
	}
}

func FilePath(value utils_types.FilePath) typelog.LogType {
	return func(c *typelog.LogAtrs) {
		c.Append(slog.String("filepath", fmt.Sprintf("%v", value)))
	}
}

func Filepaths(values []utils_types.FilePath) typelog.LogType {
	return typelog.Items[utils_types.FilePath](values, "filepaths")
}
