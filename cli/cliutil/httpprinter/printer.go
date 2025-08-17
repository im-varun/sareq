package httpprinter

import (
	"fmt"
	"slices"
	"strings"

	"github.com/im-varun/sareq/cli/cliutil/httpprinter/coloring"
	"github.com/im-varun/sareq/internal/httpclient"
)

func PrintResponse(r *httpclient.Response, noColor bool, noPrettify bool) {
	respColoring := coloring.InitResponseColoring()

	if noColor {
		respColoring.DisableColors()
	}

	respColoring.Protocol("%s ", r.Protocol())

	statusCode := r.StatusCode()
	switch {
	case statusCode >= 200 && statusCode < 300:
		respColoring.StatusSuccess("%s\n", r.Status())
	default:
		respColoring.StatusFailure("%s\n", r.Status())
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
		respColoring.HeaderKey("%s: ", key)

		values := header[key]
		switch len(values) {
		case 0:
			respColoring.HeaderValue("%s\n", "empty")
		case 1:
			respColoring.HeaderValue("%s\n", values[0])
		default:
			respColoring.HeaderValue("%v\n", values)
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

	if bodyIsJSONType && !noPrettify {
		prettyBody, err := prettifyResponseBody(body)
		if err != nil {
			// body could not be prettifyed, so printing it normally
			respColoring.Body("%s\n", body)
		} else {
			respColoring.Body("%s\n", prettyBody)
		}
	} else {
		respColoring.Body("%s\n", body)
	}
}
