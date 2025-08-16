package httpclient

import (
	"io"
	"net/http"
	"strings"
)

func newRequest(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string) (*http.Request, error) {
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

	for key, value := range reqHeader {
		req.Header.Add(key, value)
	}

	return req, nil
}
