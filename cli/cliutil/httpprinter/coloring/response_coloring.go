package coloring

import (
	"github.com/fatih/color"
)

type ResponseColoring struct {
	Protocol      ColoredPrinter
	StatusSuccess ColoredPrinter
	StatusFailure ColoredPrinter
	HeaderKey     ColoredPrinter
	HeaderValue   ColoredPrinter
	Body          ColoredPrinter
}

func InitResponseColoring() *ResponseColoring {
	return &ResponseColoring{
		Protocol:      color.New(color.FgHiCyan, color.Bold).PrintfFunc(),
		StatusSuccess: color.New(color.FgHiGreen, color.Bold).PrintfFunc(),
		StatusFailure: color.New(color.FgHiRed, color.Bold).PrintfFunc(),
		HeaderKey:     color.New(color.FgHiMagenta).PrintfFunc(),
		HeaderValue:   color.New(color.FgHiWhite).PrintfFunc(),
		Body:          color.New(color.FgHiYellow).PrintfFunc(),
	}
}

func (rc *ResponseColoring) DisableColors() {
	rc.Protocol = NoColoredPrinter
	rc.StatusSuccess = NoColoredPrinter
	rc.StatusFailure = NoColoredPrinter
	rc.HeaderKey = NoColoredPrinter
	rc.HeaderValue = NoColoredPrinter
	rc.Body = NoColoredPrinter
}
