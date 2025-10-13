package httpclient

import (
	"io"
	"net/http"
	"strings"
)

// request is a wrapper around the standard http.Request type from net/http.
type request struct {
	httpRequest *http.Request // HTTP request from net/http
}

// newRequest creates and returns a new request instance using the specified
// request method, URL, body and header.
func newRequest(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string) (*request, error) {
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

	return &request{
		httpRequest: req,
	}, nil
}
