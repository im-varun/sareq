package httpclient

import (
	"fmt"
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
		return nil, fmt.Errorf("unable to create a request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to '%s': %w", reqURL, normalizeHTTPError(err))
	}
	defer resp.Body.Close()

	response, err := parseResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}

	return response, nil
}
