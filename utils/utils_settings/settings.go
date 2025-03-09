package utils_settings

import (
	"github.com/darklab8/go-utils/utils/enverant"
)

type UtilsEnvs struct {
	IsDevEnv             bool
	AreTestsRegenerating bool
	Enver                *enverant.Enverant
}

var Envs UtilsEnvs

func init() {
	GetEnvs()
}

func GetEnvs() UtilsEnvs {
	envs := enverant.NewEnverant(enverant.WithPrefix("UTILS_"), enverant.WithDescription("UTILS set of envs for lib of small stuff reusable for across any app"))

	Envs = UtilsEnvs{
		IsDevEnv:             envs.GetBool("DEV_ENV", enverant.OrBool(false), enverant.WithDesc("if u wish running smth differently when running darkstat in IDE, that's your option")),
		AreTestsRegenerating: envs.GetBool("TEST_REGENERATE", enverant.OrBool(false), enverant.WithDesc("if u wish to use current test run to regenerate unit tests, that's your option")),
		Enver:                envs,
	}
	return Envs
}
