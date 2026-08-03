package typelog

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricMsgsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "typelog_msgs_total",
		Help: "Message levels",
	}, []string{
		"level",
		"location",
		"is_printed",
	})

	CountMsgsWarns  = 0
	CountMsgsErrors = 0
	CountMsgsPanics = 0
	CountMsgsInfos  = 0
	CountMsgsDebugs = 0
)
