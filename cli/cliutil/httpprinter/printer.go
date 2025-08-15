package httpprinter

import (
	"fmt"
	"slices"
	"strings"

	"github.com/im-varun/sareq/internal/httpclient"
)

func PrintResponse(r *httpclient.Response) {
	respPrinter := newResponsePrinter()

	respPrinter.Protocol("%s ", r.Protocol())

	statusCode := r.StatusCode()
	switch {
	case statusCode >= 200 && statusCode < 300:
		respPrinter.StatusGreen("%s\n", r.Status())
	case statusCode >= 400 && statusCode < 500:
		respPrinter.StatusYellow("%s\n", r.Status())
	case statusCode >= 500 && statusCode < 600:
		respPrinter.StatusRed("%s\n", r.Status())
	default:
		fmt.Printf("%s\n", r.Status())
	}

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
		respPrinter.HeaderKey("%s: ", key)

		values := header[key]
		switch len(values) {
		case 0:
			respPrinter.HeaderValue("%s\n", "empty")
		case 1:
			respPrinter.HeaderValue("%s\n", values[0])
		default:
			respPrinter.HeaderValue("%v\n", values)
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
			respPrinter.Body("%s\n", body)
		} else {
			respPrinter.Body("%s\n", prettyBody)
		}
	} else {
		respPrinter.Body("%s\n", body)
	}
}
