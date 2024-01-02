package utils

import (
	"regexp"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/logus/logus_types"
	"github.com/darklab8/darklab_goutils/goutils/logus/utils_logus"
)

func InitRegexExpression(regex **regexp.Regexp, expression logus_types.RegExp) {
	var err error

	*regex, err = regexp.Compile(string(expression))
	utils_logus.Log.CheckFatal(err, "failed to init regex", logus.Regex(expression), logus.FilePath(GetCurrentFile()))
}
