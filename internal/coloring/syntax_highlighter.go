package coloring

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/fatih/color"
)

// SyntaxHighlighterFunc defines a function type for printing formatted strings with complex
// syntax highlighting.
type SyntaxHighlighterFunc func(typ string, format string, a ...any)

// NewSyntaxHighlighterFunc creates and returns a SyntaxHighlighterFunc that prints text with
// the specified syntax highlighting attributes.
func NewSyntaxHighlighterFunc(fmtter string, styling string, fallbackAttrs ...color.Attribute) SyntaxHighlighterFunc {
	return func(typ string, format string, a ...any) {
		fallbackColoredPrinter := NewColoredPrinterFunc(fallbackAttrs...)

		lexer := lexers.Get(typ)
		if lexer == nil {
			fallbackColoredPrinter(format, a...)
			return
		}

		formatter := formatters.Get(fmtter)
		if formatter == nil {
			fallbackColoredPrinter(format, a...)
			return
		}

		style := styles.Get(styling)
		if style == nil {
			fallbackColoredPrinter(format, a...)
			return
		}

		str := fmt.Sprintf(format, a...)

		iterator, err := lexer.Tokenise(nil, str)
		if err != nil {
			fallbackColoredPrinter(format, a...)
			return
		}

		var buf bytes.Buffer
		err = formatter.Format(&buf, style, iterator)
		if err != nil {
			fallbackColoredPrinter(format, a...)
			return
		}

		fmt.Print(buf.String())
	}
}

// NoSyntaxHighlighterFunc creates and returns a SyntaxHighlighterFunc that prints text without
// any syntax highlighting attributes.
func NoSyntaxHighlighterFunc() SyntaxHighlighterFunc {
	return func(typ string, format string, a ...any) {
		fmt.Printf(format, a...)
	}
}
