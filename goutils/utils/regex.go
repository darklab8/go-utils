package utils

import (
	"regexp"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logger"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func InitRegexExpression(regex **regexp.Regexp, expression utils_types.RegExp) {
	var err error

	*regex, err = regexp.Compile(string(expression))
	utils_logger.Log.CheckFatal(err, "failed to init regex",
		utils_logger.Regex(expression), utils_logger.FilePath(GetCurrentFile()))
}
