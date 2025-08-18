package coloring

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/fatih/color"
)

type ResponseColoring struct {
	Protocol      ColoredPrinter
	StatusSuccess ColoredPrinter
	StatusFailure ColoredPrinter
	HeaderKey     ColoredPrinter
	HeaderValue   ColoredPrinter
	Body          SyntaxHighlighter
}

func InitResponseColoring() *ResponseColoring {
	return &ResponseColoring{
		Protocol:      color.New(color.FgHiCyan, color.Bold).PrintfFunc(),
		StatusSuccess: color.New(color.FgHiGreen, color.Bold).PrintfFunc(),
		StatusFailure: color.New(color.FgHiRed, color.Bold).PrintfFunc(),
		HeaderKey:     color.New(color.FgHiMagenta).PrintfFunc(),
		HeaderValue:   color.New(color.FgHiWhite).PrintfFunc(),
		Body:          NewSyntaxHighlighterFunc(),
	}
}

func NewSyntaxHighlighterFunc() SyntaxHighlighter {
	return func(typ string, format string, a ...any) {
		baseColoring := color.New(color.FgHiYellow).PrintfFunc()

		lexer := lexers.Get(typ)
		if lexer == nil {
			baseColoring(format, a...)
			return
		}

		formatter := formatters.Get("terminal16m")
		if formatter == nil {
			baseColoring(format, a...)
			return
		}

		style := styles.Get("monokai")
		if style == nil {
			baseColoring(format, a...)
			return
		}

		str := fmt.Sprintf(format, a...)

		iterator, err := lexer.Tokenise(nil, str)
		if err != nil {
			baseColoring(format, a...)
			return
		}

		var buf bytes.Buffer
		err = formatter.Format(&buf, style, iterator)
		if err != nil {
			baseColoring(format, a...)
			return
		}

		fmt.Print(buf.String())
	}
}

func (rc *ResponseColoring) DisableColors() {
	rc.Protocol = NoColoredPrinter
	rc.StatusSuccess = NoColoredPrinter
	rc.StatusFailure = NoColoredPrinter
	rc.HeaderKey = NoColoredPrinter
	rc.HeaderValue = NoColoredPrinter
	rc.Body = NoSyntaxHighlighter
}
