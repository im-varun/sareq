package httprunner

import (
	"fmt"

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

	resp, err := client.Execute(reqMethod, reqURL, flags.ReqBody)
	if err != nil {
		return err
	}

	printResponse(resp)

	return nil
}

func printResponse(r *httpclient.Response) {
	fmt.Printf("%s %s\n", r.Protocol(), r.Status())
	fmt.Printf("Content-Length: %d\n\n", r.ContentLength())

	header := r.Header()
	for key, values := range header {
		if len(values) == 1 {
			fmt.Printf("%s: %s\n", key, values[0])
		} else {
			fmt.Printf("%s: %v\n", key, values)
		}
	}
	fmt.Println()

	fmt.Printf("%s\n", r.Body())
}
