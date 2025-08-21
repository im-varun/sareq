package coloring

import (
	"github.com/fatih/color"
)

type ResponseColoring struct {
	Protocol      ColoredPrinterFunc
	StatusSuccess ColoredPrinterFunc
	StatusFailure ColoredPrinterFunc
	HeaderKey     ColoredPrinterFunc
	HeaderValue   ColoredPrinterFunc
	Body          SyntaxHighlighterFunc
}

func InitResponseColoring() *ResponseColoring {
	return &ResponseColoring{
		Protocol:      NewColoredPrinterFunc(color.FgHiCyan, color.Bold),
		StatusSuccess: NewColoredPrinterFunc(color.FgHiGreen, color.Bold),
		StatusFailure: NewColoredPrinterFunc(color.FgHiRed, color.Bold),
		HeaderKey:     NewColoredPrinterFunc(color.FgHiMagenta),
		HeaderValue:   NewColoredPrinterFunc(color.FgHiWhite),
		Body:          NewSyntaxHighlighterFunc(color.FgHiYellow),
	}
}

func (rc *ResponseColoring) DisableColors() {
	rc.Protocol = NoColoredPrinterFunc
	rc.StatusSuccess = NoColoredPrinterFunc
	rc.StatusFailure = NoColoredPrinterFunc
	rc.HeaderKey = NoColoredPrinterFunc
	rc.HeaderValue = NoColoredPrinterFunc
	rc.Body = NoSyntaxHighlighterFunc
}
