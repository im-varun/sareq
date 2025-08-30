package httprunner

import (
	"fmt"

	"github.com/im-varun/sareq/internal/httpclient"
)

func Run(reqMethod string, reqURL string, reqBody string, reqHeader map[string]string, reqTimeout int) (*httpclient.Response, error) {
	validatedReqURL, err := httpclient.ValidateRequestURL(reqURL)
	if err != nil {
		return nil, fmt.Errorf("request URL '%s' is invalid: %w", reqURL, err)
	}

	if reqBody != "" {
		err = httpclient.ValidateRequestBody(reqBody)
		if err != nil {
			return nil, fmt.Errorf("request body '%s' is invalid: %w", reqBody, err)
		}
	}

	client := httpclient.NewClient(reqTimeout)

	resp, err := client.Execute(reqMethod, validatedReqURL, reqBody, reqHeader)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
