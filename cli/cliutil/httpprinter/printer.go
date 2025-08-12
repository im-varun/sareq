package httpprinter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/im-varun/sareq/internal/httpclient"
)

func PrintResponse(r *httpclient.Response) {
	fmt.Println(r.Protocol(), r.Status())
	fmt.Println()

	header := r.Header()
	for key, values := range header {
		if len(values) == 1 {
			fmt.Printf("%s: %s\n", key, values[0])
		} else {
			fmt.Printf("%s: %v\n", key, values)
		}
	}
	fmt.Println()

	body := strings.Trim(r.Body(), "\r\n")

	var bodyIsJSONType bool

	contentType, ok := header["Content-Type"]
	if ok {
		for _, str := range contentType {
			if strings.Contains(str, "application/json") {
				bodyIsJSONType = true
				break
			}
		}
	}

	if bodyIsJSONType {
		prettyBody, err := prettifyResponseBody(body)
		if err != nil {
			// body could not be prettifyed, so printing it normally
			fmt.Println(body)
		} else {
			fmt.Println(prettyBody)
		}
	} else {
		fmt.Println(body)
	}
}

func prettifyResponseBody(respBody string) (string, error) {
	var prettyBody bytes.Buffer

	err := json.Indent(&prettyBody, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	return prettyBody.String(), nil
}
