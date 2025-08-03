package utils_settings

import (
	"github.com/darklab8/go-utils/utils/enverant"
)

type UtilsEnvs struct {
	IsDevEnv             bool
	AreTestsRegenerating bool
	Environment          string
	VersionId            string
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
		Environment:          envs.GetStr("ENVIRONMENT", enverant.OrStr("undefined"), enverant.WithDesc("Environment like staging or production or anything else where app runs. Adds metric label")),
		VersionId:            envs.GetStr("VERSION_ID", enverant.OrStr("undefined")),
		Enver:                envs,
	}
	return Envs
}
