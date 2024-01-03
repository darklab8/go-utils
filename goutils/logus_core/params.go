package logus_core

import (
	"fmt"
	"log/slog"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func logGroupFiles() slog.Attr {
	return slog.Group("files",
		"file3", GetCallingFile(3),
		"file4", GetCallingFile(4),
	)
}

type SlogGroup struct {
	Params map[string]string
}

func (s SlogGroup) Render() slog.Attr {
	anies := []any{}
	for key, value := range s.Params {
		anies = append(anies, key)
		anies = append(anies, value)
	}

	return slog.Group("extras", anies...)
}

type SlogParam func(r *SlogGroup)

func newSlogGroup(opts ...SlogParam) slog.Attr {
	client := &SlogGroup{Params: make(map[string]string)}
	for _, opt := range opts {
		opt(client)
	}

	return (*client).Render()
}

func TestParam(value int) SlogParam {
	return func(c *SlogGroup) {
		c.Params["test_param"] = fmt.Sprintf("%d", value)
	}
}

func Expected(value any) SlogParam {
	return func(c *SlogGroup) {
		c.Params["expected"] = fmt.Sprintf("%v", value)
	}
}
func Actual(value any) SlogParam {
	return func(c *SlogGroup) {
		c.Params["actual"] = fmt.Sprintf("%v", value)
	}
}

func OptError(err error) SlogParam {
	return func(c *SlogGroup) {
		c.Params["error_msg"] = fmt.Sprintf("%v", err)
		c.Params["error_type"] = fmt.Sprintf("%T", err)
	}
}

func FilePath(value utils_types.FilePath) SlogParam {
	return func(c *SlogGroup) {
		c.Params["filepath"] = fmt.Sprintf("%v", value)
	}
}

func Items[T any](value []T, item_name string) SlogParam {
	return func(c *SlogGroup) {
		sliced_string := fmt.Sprintf("%v", value)
		if len(sliced_string) > 300 {
			sliced_string = sliced_string[:300] + "...sliced string"
		}
		c.Params[item_name] = sliced_string
		c.Params[fmt.Sprintf("%s_len", item_name)] = fmt.Sprintf("%d", len(value))
	}
}

func Records[T any](value []T) SlogParam {
	return Items[T](value, "records")
}

func Args(value []string) SlogParam {
	return Items[string](value, "args")
}

func Body(value []byte) SlogParam {
	return func(c *SlogGroup) {
		c.Params["body"] = string(value)
	}
}

func ErrorMsg(value string) SlogParam {
	return func(c *SlogGroup) {
		c.Params["error_message"] = string(value)
	}
}
