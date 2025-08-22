package httpprinter

import (
	"bytes"
	"encoding/json"

	"github.com/fatih/color"
	"github.com/im-varun/sareq/internal/coloring"
)

type responseColoring struct {
	protocol      coloring.ColoredPrinterFunc
	statusSuccess coloring.ColoredPrinterFunc
	statusFailure coloring.ColoredPrinterFunc
	headerKey     coloring.ColoredPrinterFunc
	headerValue   coloring.ColoredPrinterFunc
	body          coloring.SyntaxHighlighterFunc
}

func initResponseColoring() *responseColoring {
	return &responseColoring{
		protocol:      coloring.NewColoredPrinterFunc(color.FgHiCyan, color.Bold),
		statusSuccess: coloring.NewColoredPrinterFunc(color.FgHiGreen, color.Bold),
		statusFailure: coloring.NewColoredPrinterFunc(color.FgHiRed, color.Bold),
		headerKey:     coloring.NewColoredPrinterFunc(color.FgHiMagenta),
		headerValue:   coloring.NewColoredPrinterFunc(color.FgHiWhite),
		body:          coloring.NewSyntaxHighlighterFunc(color.FgHiYellow),
	}
}

func (rc *responseColoring) disableColors() {
	rc.protocol = coloring.NoColoredPrinterFunc()
	rc.statusSuccess = coloring.NoColoredPrinterFunc()
	rc.statusFailure = coloring.NoColoredPrinterFunc()
	rc.headerKey = coloring.NoColoredPrinterFunc()
	rc.headerValue = coloring.NoColoredPrinterFunc()
	rc.body = coloring.NoSyntaxHighlighterFunc()
}

func prettifyResponseBody(respBody string) (string, error) {
	var prettyBodyBuf bytes.Buffer

	err := json.Indent(&prettyBodyBuf, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	prettyBody := prettyBodyBuf.String()

	return prettyBody, nil
}
