package httpclient

import (
	"io"
	"net/http"
)

type Response struct {
	status   string
	protocol string
	header   map[string][]string
	body     string
}

func parseResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		status:   resp.Status,
		protocol: resp.Proto,
		header:   resp.Header,
		body:     string(body),
	}, nil
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) Protocol() string {
	return r.protocol
}

func (r *Response) Header() map[string][]string {
	return r.header
}

func (r *Response) Body() string {
	return r.body
}
