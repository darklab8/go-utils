package logus_core

import (
	"os"
)

var EnvTurnJSON bool = os.Getenv("GO_LOG_JSON") == "true"

var EnvTurnFileShowing bool = os.Getenv("GO_LOG_FILE_SHOWING") == "true"
