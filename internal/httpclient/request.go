package httpclient

import (
	"io"
	"net/http"
	"strings"
)

func newRequest(reqMethod string, reqURL string, reqBody string) (*http.Request, error) {
	method := strings.ToUpper(reqMethod)

	url := reqURL

	var body io.Reader
	if reqBody != "" {
		body = strings.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
