package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

type responseData struct {
	status        string
	body          string
	contentLength int64
}

func newResponseData(resp *http.Response) (*responseData, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &responseData{
		status:        resp.Status,
		body:          string(body),
		contentLength: resp.ContentLength,
	}, nil
}

func (rd *responseData) print() {
	fmt.Println("Status:", rd.status)
	fmt.Println(rd.body)
}
