package httprunner

import (
	"github.com/im-varun/sareq/cli/flags"
	"github.com/im-varun/sareq/internal/httpclient"
)

func Run(reqMethod string, reqURL string) (*httpclient.Response, error) {
	reqURL, err := httpclient.ValidateRequestURL(reqURL)
	if err != nil {
		return nil, err
	}

	if flags.ReqBody != "" {
		err = httpclient.ValidateRequestBody(flags.ReqBody)
		if err != nil {
			return nil, err
		}
	}

	client := httpclient.NewClient(flags.ReqTimeout)

	resp, err := client.Execute(reqMethod, reqURL, flags.ReqBody)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
