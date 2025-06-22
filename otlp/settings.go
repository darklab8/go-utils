package otlp

import (
	"github.com/darklab8/go-utils/utils/enverant"
	"github.com/darklab8/go-utils/utils/utils_settings"
)

type OtlpEnvVars struct {
	utils_settings.UtilsEnvs
	HttpOn          bool
	MetricsInterval int
}

var Env OtlpEnvVars

func GetEnvs() OtlpEnvVars {
	envs := enverant.NewEnverant(enverant.WithPrefix("OTLP_"), enverant.WithDescription("OTLP related env vars"))

	Env = OtlpEnvVars{
		UtilsEnvs:       utils_settings.GetEnvs(),
		HttpOn:          envs.GetBool("HTTP_ON", enverant.WithDesc("start submit to http endpoint")),
		MetricsInterval: envs.GetIntOr("OTLP_METRICS_INTERVAL_SECS", 60, enverant.WithDesc("interval in seconds to submit metrics, default 1m")),
	}
	return Env
}

func init() {
	Env = GetEnvs()
}
