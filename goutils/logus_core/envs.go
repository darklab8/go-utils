package logus_core

import (
	"os"

	"github.com/darklab8/darklab_goutils/goutils/logus_core/logus_types"
)

var EnvTurnJSON logus_types.EnableJsonFormat = logus_types.EnableJsonFormat(os.Getenv("GO_LOG_JSON") == "true")

var EnvTurnFileShowing bool = os.Getenv("GO_LOG_FILE_SHOWING") == "true"
