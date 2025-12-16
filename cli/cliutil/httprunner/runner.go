package httprunner

import (
	"github.com/im-varun/sareq/internal/httpclient"
)

// Run runs an HTTP request using the specified request method, URL, body, header and timeout.
func Run(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string, reqTimeout int) (*httpclient.Response, error) {
	client := httpclient.NewClient(reqTimeout)
	resp, err := client.Execute(reqMethod, reqURL, reqBody, reqHeader)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
