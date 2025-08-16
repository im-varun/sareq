package httprunner

import (
	"github.com/im-varun/sareq/internal/httpclient"
)

func Run(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string, reqTimeout int) (*httpclient.Response, error) {
	reqURL, err := httpclient.ValidateRequestURL(reqURL)
	if err != nil {
		return nil, err
	}

	if reqBody != "" {
		err = httpclient.ValidateRequestBody(reqBody)
		if err != nil {
			return nil, err
		}
	}

	client := httpclient.NewClient(reqTimeout)

	resp, err := client.Execute(reqMethod, reqURL, reqBody, reqHeader)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
