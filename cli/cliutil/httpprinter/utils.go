package httpprinter

import (
	"bytes"
	"encoding/json"

	"github.com/fatih/color"
	"github.com/im-varun/sareq/internal/coloring"
)

// responseColoring is a collection of custom coloring functions used to print different parts
// of an HTTP response.
type responseColoring struct {
	protocol      coloring.ColoredPrinterFunc
	statusSuccess coloring.ColoredPrinterFunc
	statusFailure coloring.ColoredPrinterFunc
	headerKey     coloring.ColoredPrinterFunc
	headerValue   coloring.ColoredPrinterFunc
	body          coloring.SyntaxHighlighterFunc
}

// initResponseColoring creates and returns a new responseColoring instance containing a
// collection of custom coloring functions used for printing different parts of an HTTP
// response.
func initResponseColoring() *responseColoring {
	return &responseColoring{
		protocol:      coloring.NewColoredPrinterFunc(color.FgHiCyan, color.Bold),
		statusSuccess: coloring.NewColoredPrinterFunc(color.FgHiGreen, color.Bold),
		statusFailure: coloring.NewColoredPrinterFunc(color.FgHiRed, color.Bold),
		headerKey:     coloring.NewColoredPrinterFunc(color.FgHiMagenta),
		headerValue:   coloring.NewColoredPrinterFunc(color.FgHiWhite),
		body:          coloring.NewSyntaxHighlighterFunc("terminal16m", "dracula", color.FgHiYellow),
	}
}

// disable disables the coloring functions contained in the responseColoring instance.
func (rc *responseColoring) disable() {
	rc.protocol = coloring.NoColoredPrinterFunc()
	rc.statusSuccess = coloring.NoColoredPrinterFunc()
	rc.statusFailure = coloring.NoColoredPrinterFunc()
	rc.headerKey = coloring.NoColoredPrinterFunc()
	rc.headerValue = coloring.NoColoredPrinterFunc()
	rc.body = coloring.NoSyntaxHighlighterFunc()
}

// prettifyResponseBody prettifies a JSON response body.
func prettifyResponseBody(respBody string) (string, error) {
	var prettyBodyBuf bytes.Buffer

	err := json.Indent(&prettyBodyBuf, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	prettyBody := prettyBodyBuf.String()

	return prettyBody, nil
}
