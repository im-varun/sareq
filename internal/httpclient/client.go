package httpclient

import (
	"net/http"
	"time"
)

const (
	DefaultTimeoutSeconds int = 10
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout int) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

func (c *Client) Do(reqMethod string, reqURL string, reqBody string) error {
	reqConfig := newRequestConfig(reqMethod, reqURL, reqBody)

	req, err := reqConfig.buildRequest()
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := newResponseData(resp)
	if err != nil {
		return err
	}

	respData.print()

	return nil
}
