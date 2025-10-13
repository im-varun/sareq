package httpclient

import (
	"io"
	"net/http"
)

// Response is a wrapper around the standard http.Response type from net/http.
type Response struct {
	status     string              // HTTP response status (e.g. "200 OK")
	statusCode int                 // HTTP response status code (e.g. 200)
	protocol   string              // HTTP response protocol (e.g. "HTTP/1.0")
	header     map[string][]string // HTTP response header
	body       string              // HTTP response body
}

// parseResponse reads and parses an http.Response instance into a Response
// type struct.
func parseResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		status:     resp.Status,
		statusCode: resp.StatusCode,
		protocol:   resp.Proto,
		header:     resp.Header,
		body:       string(body),
	}, nil
}

// Status returns the status of the HTTP response.
func (r *Response) Status() string {
	return r.status
}

// StatusCode returns the status code of the HTTP response.
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Protocol returns the protocol of the HTTP response.
func (r *Response) Protocol() string {
	return r.protocol
}

// Header returns the header of the HTTP response.
func (r *Response) Header() map[string][]string {
	headerCopy := make(map[string][]string, len(r.header))

	for key, values := range r.header {
		valuesCopy := make([]string, len(values))
		copy(valuesCopy, values)
		headerCopy[key] = valuesCopy
	}

	return headerCopy
}

// Body returns the body of the HTTP response.
func (r *Response) Body() string {
	return r.body
}
