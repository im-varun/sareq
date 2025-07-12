package httpclient

import (
	"io"
	"net/http"
	"strings"
)

type RequestConfig struct {
	method string
	url    string
	body   string
}

func NewRequestConfig(method string, url string, body string) *RequestConfig {
	return &RequestConfig{
		method: method,
		url:    url,
		body:   body,
	}
}

func (rc *RequestConfig) Build() (*http.Request, error) {
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
