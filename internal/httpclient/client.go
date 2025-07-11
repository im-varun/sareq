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
