package httpprinter

import (
	"fmt"
	"mime"
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

	keys := make([]string, 0, len(header))
	for key := range header {
		keys = append(keys, key)
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

	contentType := header["Content-Type"]
	mediaType, _, err := mime.ParseMediaType(contentType[0])
	if err != nil {
		mediaType = "text/plain"
	}

	bodyTextType := strings.Split(mediaType, "/")[1]

	if mediaType == "application/json" && !noPrettify {
		prettyBody, err := prettifyResponseBody(body)
		if err != nil {
			respColoring.Body(bodyTextType, "%s\n", body)
		} else {
			respColoring.Body(bodyTextType, "%s\n", prettyBody)
		}
	} else {
		respColoring.Body(bodyTextType, "%s\n", body)
	}
}
