package httpclient

import (
	"io"
	"net/http"
)

type Response struct {
	status        string
	body          string
	contentLength int64
}

func parseResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		status:        resp.Status,
		body:          string(body),
		contentLength: resp.ContentLength,
	}, nil
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) Body() string {
	return r.body
}

func (r *Response) ContentLength() int64 {
	return r.contentLength
}
