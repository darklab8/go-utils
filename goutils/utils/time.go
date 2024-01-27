package utils

import (
	"fmt"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logus"
	"github.com/darklab8/logusgo/logcore"
)

type timeMeasurer struct {
	msg          string
	ops          []logcore.SlogParam
	time_started time.Time
}

func NewTimeMeasure(msg string, ops ...logcore.SlogParam) *timeMeasurer {
	return &timeMeasurer{
		msg:          msg,
		ops:          ops,
		time_started: time.Now(),
	}
}

func (t *timeMeasurer) Close() {
	utils_logus.Log.Debug(fmt.Sprintf("time_measure %v | %s", time.Since(t.time_started), t.msg), t.ops...)
}

func TimeMeasure(callback func(), msg string, ops ...logcore.SlogParam) {
	time_started := NewTimeMeasure(msg, ops...)
	defer time_started.Close()
	callback()
}
