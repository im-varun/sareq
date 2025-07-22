package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

type responseData struct {
	respStatus        string
	respBody          string
	respContentLength int64
}

func newResponseData(resp *http.Response) (*responseData, error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &responseData{
		respStatus:        resp.Status,
		respBody:          string(respBody),
		respContentLength: resp.ContentLength,
	}, nil
}

func (rd *responseData) print() {
	fmt.Println("Status:", rd.respStatus)
	fmt.Println(rd.respBody)
}
