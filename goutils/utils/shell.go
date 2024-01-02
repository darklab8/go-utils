package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/logus/utils_logus"
)

func ShellRunArgs(program string, args ...string) {
	utils_logus.Log.Debug(fmt.Sprintf("OK attempting to run: %s", program), logus.Args(args))
	executable, _ := exec.LookPath(program)

	args = append([]string{""}, args...)
	command := exec.Cmd{
		Path:   executable,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err := command.Run()

	utils_logus.Log.CheckFatal(err, "failed to run shell command")
}
