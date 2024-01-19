package logus_core

import (
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
		c.Params["nested"] = map[string]any{
			"smth":   "abc",
			"number": 123,
		}
	}
}

type Smth struct {
	Value1  string
	Number1 int
}

func NestedStructTest(value string) SlogParam {
	return func(c *SlogGroup) {
		c.Params["nested"] = StructToMap(Smth{Value1: "123", Number1: 4})
		c.Params["not_nested"] = 345
	}
}

func TestNested(t *testing.T) {
	logger := NewLogger(LEVEL_DEBUG, logus_types.EnableJsonFormat(true), logus_types.EnableFileShowing(false))

	logger.Debug("123", NestedParam("abc"))
	logger.Debug("456", NestedStructTest("abc"))
}
