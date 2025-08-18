package coloring

import "fmt"

type ColoredPrinter func(format string, a ...any)

func NoColoredPrinter(format string, a ...any) {
	fmt.Printf(format, a...)
}

type SyntaxHighlighter func(typ string, format string, a ...any)

func NoSyntaxHighlighter(typ string, format string, a ...any) {
	fmt.Printf(format, a...)
}
