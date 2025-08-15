package httpprinter

import (
	"bytes"
	"encoding/json"

	"github.com/fatih/color"
)

type Printer func(format string, a ...interface{})

type ResponsePrinter struct {
	Protocol     Printer
	StatusGreen  Printer
	StatusYellow Printer
	StatusRed    Printer
	HeaderKey    Printer
	HeaderValue  Printer
	Body         Printer
}

func newResponsePrinter() *ResponsePrinter {
	return &ResponsePrinter{
		Protocol:     color.New(color.FgCyan, color.Bold).PrintfFunc(),
		StatusGreen:  color.New(color.FgGreen, color.Bold).PrintfFunc(),
		StatusYellow: color.New(color.FgYellow, color.Bold).PrintfFunc(),
		StatusRed:    color.New(color.FgRed, color.Bold).PrintfFunc(),
		HeaderKey:    color.New(color.FgMagenta).PrintfFunc(),
		HeaderValue:  color.New(color.FgHiWhite).PrintfFunc(),
		Body:         color.New(color.FgYellow).PrintfFunc(),
	}
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
