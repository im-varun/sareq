package coloring

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/fatih/color"
)

type SyntaxHighlighterFunc func(typ string, format string, a ...any)

func NewSyntaxHighlighterFunc(baseColor color.Attribute) SyntaxHighlighterFunc {
	return func(typ string, format string, a ...any) {
		baseColoredPrinter := NewColoredPrinterFunc(baseColor)

		lexer := lexers.Get(typ)
		if lexer == nil {
			baseColoredPrinter(format, a...)
			return
		}

		formatter := formatters.Get("terminal16m")
		if formatter == nil {
			baseColoredPrinter(format, a...)
			return
		}

		style := styles.Get("dracula")
		if style == nil {
			baseColoredPrinter(format, a...)
			return
		}

		str := fmt.Sprintf(format, a...)

		iterator, err := lexer.Tokenise(nil, str)
		if err != nil {
			baseColoredPrinter(format, a...)
			return
		}

		var buf bytes.Buffer
		err = formatter.Format(&buf, style, iterator)
		if err != nil {
			baseColoredPrinter(format, a...)
			return
		}

		fmt.Print(buf.String())
	}
}

func NoSyntaxHighlighterFunc(typ string, format string, a ...any) {
	fmt.Printf(format, a...)
}
