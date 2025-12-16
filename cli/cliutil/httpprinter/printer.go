package httpprinter

import (
	"fmt"
	"mime"
	"slices"
	"strings"

	"github.com/fatih/color"
	"github.com/im-varun/sareq/internal/coloring"
	"github.com/im-varun/sareq/internal/httpclient"
)

// Print prints the specified HTTP response with/without coloring and with/without JSON
// prettification formatting.
func Print(resp *httpclient.Response, respNoColor bool, respNoPrettify bool) {
	var printProtocol coloring.ColoredPrinterFunc
	var printStatusSuccess coloring.ColoredPrinterFunc
	var printStatusNotSuccess coloring.ColoredPrinterFunc
	var printHeaderKey coloring.ColoredPrinterFunc
	var printHeaderValue coloring.ColoredPrinterFunc
	var printBody coloring.SyntaxHighlighterFunc

	if !respNoColor {
		printProtocol = coloring.NewColoredPrinterFunc(color.FgHiCyan, color.Bold)
		printStatusSuccess = coloring.NewColoredPrinterFunc(color.FgHiGreen, color.Bold)
		printStatusNotSuccess = coloring.NewColoredPrinterFunc(color.FgHiRed, color.Bold)
		printHeaderKey = coloring.NewColoredPrinterFunc(color.FgHiMagenta)
		printHeaderValue = coloring.NewColoredPrinterFunc(color.FgHiWhite)
		printBody = coloring.NewSyntaxHighlighterFunc("terminal16m", "dracula", color.FgHiYellow)
	} else {
		printProtocol = coloring.NoColoredPrinterFunc()
		printStatusSuccess = coloring.NoColoredPrinterFunc()
		printStatusNotSuccess = coloring.NoColoredPrinterFunc()
		printHeaderKey = coloring.NoColoredPrinterFunc()
		printHeaderValue = coloring.NoColoredPrinterFunc()
		printBody = coloring.NoSyntaxHighlighterFunc()
	}

	printProtocol("%s ", resp.Protocol())

	statusCode := resp.StatusCode()
	if statusCode >= 200 && statusCode < 300 {
		printStatusSuccess("%s\n", resp.Status())
	} else {
		printStatusNotSuccess("%s\n", resp.Status())
	}
	fmt.Println()

	header := resp.Header()

	keys := make([]string, 0, len(header))
	for key := range header {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	for _, key := range keys {
		printHeaderKey("%s: ", key)

		values := header[key]
		switch len(values) {
		case 0:
			printHeaderValue("%s\n", "empty")
		case 1:
			printHeaderValue("%s\n", values[0])
		default:
			printHeaderValue("%v\n", values)
		}
	}
	fmt.Println()

	contentType := header["Content-Type"]
	mediaType, _, err := mime.ParseMediaType(contentType[0])
	if err != nil {
		mediaType = "text/plain"
	}

	if isTextualMediaType(mediaType) {
		body := strings.Trim(resp.Body(), "\r\n")
		bodyType := strings.Split(mediaType, "/")[1]

		if mediaType == "application/json" && !respNoPrettify {
			prettyBody, err := prettifyResponseBody(body)
			if err != nil {
				printBody(bodyType, "%s\n", body)
			} else {
				printBody(bodyType, "%s\n", prettyBody)
			}
		} else {
			printBody(bodyType, "%s\n", body)
		}
	} else {
		fmt.Println("[binary data cannot be displayed]")
	}
}
