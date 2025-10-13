package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

// Client is a wrapper around the standard http.Client type from net/http.
// It allows custom configuration of the HTTP client settings.
type Client struct {
	httpClient *http.Client // HTTP client from net/http
}

// NewClient creates and returns a new Client instance with the specified timeout
// (in seconds).
func NewClient(timeout int) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// Execute executes the HTTP request with the specified request method, URL, body
// and header.
func (c *Client) Execute(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string) (*Response, error) {
	req, err := newRequest(reqMethod, reqURL, reqBody, reqHeader)
	if err != nil {
		return nil, fmt.Errorf("unable to create a request: %w", err)
	}

	response, err := c.httpClient.Do(req.httpRequest)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to '%s': %w", reqURL, normalizeNetworkError(err))
	}
	defer response.Body.Close()

	resp, err := parseResponse(response)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}

	return resp, nil
}
