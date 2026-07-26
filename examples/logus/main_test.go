package logus

import (
	"context"
	"errors"
	"testing"
)

func TestLogus(t *testing.T) {

	Log.Info("smth written for test")

	Log.Infoln("smth written for test2")
	Log.InfoCtx(context.Background(), "smth written for test3")

	err := errors.New("smth error")
	Log.CheckError(err, "abc")
	Log.CheckErrorln(err, "abc")
	Log.CheckErrorCtx(context.Background(), err, "abc")
}
