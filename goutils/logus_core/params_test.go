package logus_core

import (
	"log/slog"
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/logus_core/logus_types"
)

func TestSlogging(t *testing.T) {

	logger := NewLogger(LEVEL_DEBUG, logus_types.EnableJsonFormat(false), logus_types.EnableFileShowing(false))
	logger.Debug("123")

	logger.Debug("123", TestParam(456))
}

func NestedParam(value string) SlogParam {
	return func(c *SlogGroup) {
		c.Attrs = append(c.Attrs,
			slog.Group("nested",
				slog.Int("number", 1),
				slog.String("smth", value),
			),
		)
	}
}

func TestNested(t *testing.T) {
	logger := NewLogger(LEVEL_DEBUG, logus_types.EnableJsonFormat(true), logus_types.EnableFileShowing(false))
	logger.Debug("123", NestedParam("abc"))
}
