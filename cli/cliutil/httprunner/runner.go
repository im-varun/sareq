package httprunner

import (
	"github.com/im-varun/sareq/cli/flags"
	"github.com/im-varun/sareq/internal/httpclient"
)

func Run(reqMethod string, reqURL string) error {
	reqURL, err := httpclient.ValidateRequestURL(reqURL)
	if err != nil {
		return err
	}

	if flags.ReqBody != "" {
		err = httpclient.ValidateRequestBody(flags.ReqBody)
		if err != nil {
			return err
		}
	}

	client := httpclient.NewClient(flags.ReqTimeout)

	_, err = client.Execute(reqMethod, reqURL, flags.ReqBody)

	return err
}
