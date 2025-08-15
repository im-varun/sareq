package httpprinter

import (
	"fmt"
	"slices"
	"strings"

	"github.com/im-varun/sareq/internal/httpclient"
)

func PrintResponse(r *httpclient.Response) {
	fmt.Println(r.Protocol(), r.Status())
	fmt.Println()

	header := r.Header()

	keys := make([]string, len(header))

	i := 0
	for key := range header {
		keys[i] = key
		i++
	}
	slices.Sort(keys)

	for _, key := range keys {
		values := header[key]
		switch len(values) {
		case 0:
			fmt.Printf("%s: empty\n", key)
		case 1:
			fmt.Printf("%s: %s\n", key, values[0])
		default:
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
