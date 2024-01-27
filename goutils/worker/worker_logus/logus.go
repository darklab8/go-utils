package worker_logus

import (
	"strconv"

	"github.com/darklab8/darklab_goutils/goutils/worker/worker_types"
	"github.com/darklab8/logusgo/logcore"
)

var Log *logcore.Logger = logcore.NewLogger("worker")

func WorkerID(value worker_types.WorkerID) logcore.SlogParam {
	return func(c *logcore.SlogGroup) {
		c.Params["worker_id"] = strconv.Itoa(int(value))
	}
}

func TaskID(value worker_types.TaskID) logcore.SlogParam {
	return func(c *logcore.SlogGroup) {
		c.Params["task_id"] = strconv.Itoa(int(value))
	}
}

func StatusCode(value worker_types.TaskStatusCode) logcore.SlogParam {
	return func(c *logcore.SlogGroup) {
		c.Params["status_code"] = strconv.Itoa(int(value))
	}
}
