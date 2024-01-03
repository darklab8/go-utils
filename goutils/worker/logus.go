package worker

import (
	"fmt"
	"strings"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils"
)

func LogusStatusCodes(tasks []ITask) logus_core.SlogParam {
	str_status_codes := utils.CompL(tasks, func(x ITask) string { return fmt.Sprintf("%d", x.GetStatusCode()) })
	return func(c *logus_core.SlogGroup) {
		c.Params["status_codes"] = strings.Join(str_status_codes, ",")
	}
}
