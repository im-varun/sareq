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

func (c *Client) Execute(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string) (*Response, error) {
	req, err := newRequest(reqMethod, reqURL, reqBody, reqHeader)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
