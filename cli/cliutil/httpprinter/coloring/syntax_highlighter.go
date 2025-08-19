package coloring

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/fatih/color"
)

type SyntaxHighlighter func(typ string, format string, a ...any)

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

		style := styles.Get("dracula")
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

func NoSyntaxHighlighter(typ string, format string, a ...any) {
	fmt.Printf(format, a...)
}
