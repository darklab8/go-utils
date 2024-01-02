package worker_logus

import (
	"strconv"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/worker/worker_types"
)

func WorkerID(value worker_types.WorkerID) logus.SlogParam {
	return func(c *logus.SlogGroup) {
		c.Params["worker_id"] = strconv.Itoa(int(value))
	}
}

func TaskID(value worker_types.TaskID) logus.SlogParam {
	return func(c *logus.SlogGroup) {
		c.Params["task_id"] = strconv.Itoa(int(value))
	}
}

func StatusCode(value worker_types.TaskStatusCode) logus.SlogParam {
	return func(c *logus.SlogGroup) {
		c.Params["status_code"] = strconv.Itoa(int(value))
	}
}
