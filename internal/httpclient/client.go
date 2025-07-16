package httpclient

import (
	"net/http"
	"time"
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

func (c *Client) Do(method string, url string, body string) error {
	reqConfig := newRequestConfig(method, url, body)
	req, err := reqConfig.build()
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
