package worker_tests

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/worker"
	"github.com/darklab8/darklab_goutils/goutils/worker/worker_logus"
	"github.com/darklab8/darklab_goutils/goutils/worker/worker_types"
	"github.com/darklab8/go-typelog/typelog"

	"github.com/stretchr/testify/assert"
)

// ======================
// Test Example

type TaskTest struct {
	*worker.Task

	// any desired arbitary data
	result worker_types.TaskID
}

func NewTaskTest(id worker_types.TaskID) *TaskTest {
	return &TaskTest{Task: worker.NewTask(id)}
}

func (data *TaskTest) RunTask(worker_id worker_types.WorkerID) error {
	worker_logus.Log.Debug("task test started", worker_logus.WorkerID(worker_id), worker_logus.TaskID(data.GetID()))
	time.Sleep(time.Second * time.Duration(data.GetID()))
	data.result = data.GetID() * 1
	worker_logus.Log.Debug("task test finished", worker_logus.WorkerID(worker_id), worker_logus.TaskID(data.GetID()))
	return nil
}

func TaskResult(value worker_types.TaskID) typelog.LogType {
	return func(c *typelog.LogAtrs) {
		c.Append(slog.Int("task_result", int(value)))
	}
}

func TestWorkerTemp(t *testing.T) {
	tasks := []worker.ITask{}
	for task_id := 1; task_id <= 3; task_id++ {
		tasks = append(tasks, NewTaskTest(worker_types.TaskID(task_id)))
	}

	worker.RunTasksInTempPool(
		tasks,
		worker.WithAllowFailedTasks(),
		worker.WithDisableParallelism(false),
	)

	done_count := 0
	failed_count := 0
	for task_number, task := range tasks {
		worker_logus.Log.Debug(fmt.Sprintf("task.Done=%t", task.IsDone()), worker_logus.TaskID(worker_types.TaskID(task_number)))
		if task.IsDone() {
			done_count += 1
		} else {
			failed_count += 1
		}
	}
	assert.GreaterOrEqual(t, done_count, 3, "expected finding done tasks")
	assert.LessOrEqual(t, failed_count, 3, "expected finding failed tasks because of time sleep greater than timeout")
}
