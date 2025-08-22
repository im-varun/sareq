package httpprinter

import (
	"fmt"
	"mime"
	"slices"
	"strings"

	"github.com/im-varun/sareq/internal/httpclient"
)

func Print(resp *httpclient.Response, respNoColor bool, respNoPrettify bool) {
	respColoring := initResponseColoring()

	if respNoColor {
		respColoring.disableColors()
	}

	respColoring.protocol("%s ", resp.Protocol())

	statusCode := resp.StatusCode()
	switch {
	case statusCode >= 200 && statusCode < 300:
		respColoring.statusSuccess("%s\n", resp.Status())
	default:
		respColoring.statusFailure("%s\n", resp.Status())
	}

	fmt.Println()

	header := resp.Header()

	keys := make([]string, 0, len(header))
	for key := range header {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		respColoring.headerKey("%s: ", key)

		values := header[key]
		switch len(values) {
		case 0:
			respColoring.headerValue("%s\n", "empty")
		case 1:
			respColoring.headerValue("%s\n", values[0])
		default:
			respColoring.headerValue("%v\n", values)
		}
	}

	fmt.Println()

	body := strings.Trim(resp.Body(), "\r\n")

	contentType := header["Content-Type"]
	mediaType, _, err := mime.ParseMediaType(contentType[0])
	if err != nil {
		mediaType = "text/plain"
	}

	bodyType := strings.Split(mediaType, "/")[1]

	if mediaType == "application/json" && !respNoPrettify {
		prettyBody, err := prettifyResponseBody(body)
		if err != nil {
			respColoring.body(bodyType, "%s\n", body)
		} else {
			respColoring.body(bodyType, "%s\n", prettyBody)
		}
	} else {
		respColoring.body(bodyType, "%s\n", body)
	}
}
