package utils

import (
	"os"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logger"
)

func RegenerativeTest(callback func() error) error {
	if os.Getenv("DARK_TEST_REGENERATE") != "true" {
		utils_logger.Log.Debug("Skipping test data regenerative code")
		return nil
	}

	return callback()
}
