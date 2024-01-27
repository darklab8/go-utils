package logus_core

import (
	"testing"
	"time"
)

func TestSlogging(t *testing.T) {

	logger := NewLogger(WithLogLevel(LEVEL_DEBUG))
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
	logger := NewLogger(WithLogLevel(LEVEL_DEBUG), WithJsonFormat(true))

	logger.Debug("123", NestedParam("abc"))
	logger.Debug("456", NestedStructTest("abc"))
}

func TestCopyingLoggers(t *testing.T) {
	logger := NewLogger(WithLogLevel(LEVEL_DEBUG), WithJsonFormat(true))

	logger1 := logger.WithFields(String("smth", "123"))
	logger2 := logger1.WithFields(Int("smth2", 2))
	logger3 := logger2.WithFields(Time("smth3", time.Now()))

	logger1.Info("logger1 printed")
	logger2.Info("logger2 printed")
	logger3.Info("logger3 printed")
}
