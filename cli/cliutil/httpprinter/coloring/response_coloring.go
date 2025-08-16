package coloring

import (
	"github.com/fatih/color"
)

type ResponseColoring struct {
	Protocol     ColoredPrinter
	StatusGreen  ColoredPrinter
	StatusYellow ColoredPrinter
	StatusRed    ColoredPrinter
	HeaderKey    ColoredPrinter
	HeaderValue  ColoredPrinter
	Body         ColoredPrinter
}

func InitResponseColoring() *ResponseColoring {
	return &ResponseColoring{
		Protocol:     color.New(color.FgCyan, color.Bold).PrintfFunc(),
		StatusGreen:  color.New(color.FgGreen, color.Bold).PrintfFunc(),
		StatusYellow: color.New(color.FgYellow, color.Bold).PrintfFunc(),
		StatusRed:    color.New(color.FgRed, color.Bold).PrintfFunc(),
		HeaderKey:    color.New(color.FgMagenta).PrintfFunc(),
		HeaderValue:  color.New(color.FgHiWhite).PrintfFunc(),
		Body:         color.New(color.FgYellow).PrintfFunc(),
	}
}

func (rc *ResponseColoring) DisableColors() {
	rc.Protocol = NoColorPrinter
	rc.StatusGreen = NoColorPrinter
	rc.StatusYellow = NoColorPrinter
	rc.StatusRed = NoColorPrinter
	rc.HeaderKey = NoColorPrinter
	rc.HeaderValue = NoColorPrinter
	rc.Body = NoColorPrinter
}
