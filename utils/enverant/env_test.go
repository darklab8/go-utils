package enverant

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type EnvConf struct {
	SomeStr   string
	SomeInt   int
	SomeBool  bool
	Undefined string
}

func GetEnvs(environ *Enverant) EnvConf {
	return EnvConf{
		SomeStr:   environ.GetStr("SOME_STR"),
		SomeInt:   environ.GetInt("SOME_INT"),
		SomeBool:  environ.GetBool("SOME_BOOL", OrBool(false)),
		Undefined: environ.GetStr("UNDEFINED"),
	}
}

func TestReading(t *testing.T) {

	fmt.Println(os.Getwd())
	environ := NewEnverant(WithEnvFile(filepath.Join("testdata", "env.json")))
	envs := GetEnvs(environ)
	fmt.Println(envs)

	fmt.Println(environ.GetStr("WORKDIR"))

	environ_validator := environ.GetValidating()
	_ = environ_validator
	// GetEnvs(environ_validator)
}
