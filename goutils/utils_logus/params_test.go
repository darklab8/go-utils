package utils_logus

import (
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/utils_logus/logus_types"
)

func TestSlogging(t *testing.T) {

	logger := NewLogger(LEVEL_DEBUG, logus_types.EnableJsonFormat(false), logus_types.EnableFileShowing(false))
	logger.Debug("123")

	logger.Debug("123", TestParam(456))
}
