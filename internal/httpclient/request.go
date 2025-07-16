package httpclient

import (
	"io"
	"net/http"
	"strings"
)

type requestConfig struct {
	method string
	url    string
	body   string
}

func newRequestConfig(method string, url string, body string) *requestConfig {
	return &requestConfig{
		method: method,
		url:    url,
		body:   body,
	}
}

func (rc *requestConfig) buildRequest() (*http.Request, error) {
	method := strings.ToUpper(rc.method)

	var body io.Reader
	if rc.body != "" {
		body = strings.NewReader(rc.body)
	}

	req, err := http.NewRequest(method, rc.url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
