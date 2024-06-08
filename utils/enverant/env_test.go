package enverant

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type EnvConf struct {
	SomeStr  string
	SomeInt  int
	SomeBool bool
}

func TestReading(t *testing.T) {

	fmt.Println(os.Getwd())
	environ := NewEnverant(WithEnvFile(filepath.Join("testdata", "env.json")))
	envs := EnvConf{
		SomeStr:  environ.GetStr("SOME_STR"),
		SomeInt:  environ.GetInt("SOME_INT"),
		SomeBool: environ.GetBoolOr("SOME_BOOL", false),
	}
	fmt.Println(envs)

	fmt.Println(environ.GetStr("WORKDIR"))
}
