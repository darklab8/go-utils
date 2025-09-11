package utils_http

import (
	"net/http"

	"github.com/darklab8/go-utils/examples/logus"
	"github.com/darklab8/go-utils/utils/utils_settings"
)

var (
	Client *http.Client
)

func init() {
	Client = &http.Client{}
}

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	logus.Log.CheckPanic(err, "failed to create new Request")

	if utils_settings.Envs.UserAgent != "" {
		req.Header.Set("User-Agent", utils_settings.Envs.UserAgent)
	}

	resp, err := Client.Do(req)
	return resp, err
}
