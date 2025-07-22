package httpclient

import (
	"io"
	"net/http"
	"strings"
)

type requestConfig struct {
	reqMethod string
	reqURL    string
	reqBody   string
}

func newRequestConfig(reqMethod string, reqURL string, reqBody string) *requestConfig {
	return &requestConfig{
		reqMethod: reqMethod,
		reqURL:    reqURL,
		reqBody:   reqBody,
	}
}

func (rc *requestConfig) buildRequest() (*http.Request, error) {
	reqMethod := strings.ToUpper(rc.reqMethod)

	var reqBody io.Reader
	if rc.reqBody != "" {
		reqBody = strings.NewReader(rc.reqBody)
	}

	req, err := http.NewRequest(reqMethod, rc.reqURL, reqBody)
	if err != nil {
		return nil, err
	}

	return req, nil
}
