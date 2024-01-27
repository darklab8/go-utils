package worker

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/logusgo/logcore"
)

func LogusStatusCodes(tasks []ITask) logcore.SlogParam {
	str_status_codes := utils.CompL(tasks, func(x ITask) string { return fmt.Sprintf("%d", x.GetStatusCode()) })
	return func(c *logcore.SlogGroup) {
		c.Append(slog.String("status_codes", strings.Join(str_status_codes, ",")))
	}
}
