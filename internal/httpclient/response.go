package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

type ResponseData struct {
	status        string
	body          string
	contentLength int64
}

func NewResponseData(resp *http.Response) (*ResponseData, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseData{
		status:        resp.Status,
		body:          string(body),
		contentLength: resp.ContentLength,
	}, nil
}

func (rd *ResponseData) Print() {
	fmt.Println("Status:", rd.status)
	fmt.Println(rd.body)
}
