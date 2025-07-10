package httpclient

import (
	"net/http"
	"time"
)

type client struct {
	httpClient *http.Client
}

func NewClient(timeout int) *client {
	return &client{
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}
