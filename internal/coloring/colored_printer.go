package coloring

import (
	"fmt"

	"github.com/fatih/color"
)

// ColoredPrinterFunc defines a function type for printing formatted strings with simple
// color attributes.
type ColoredPrinterFunc func(format string, a ...any)

// NewColoredPrinterFunc creates and returns a ColoredPrinterFunc that prints text with the
// specified color attributes.
func NewColoredPrinterFunc(attrs ...color.Attribute) ColoredPrinterFunc {
	coloring := color.New(attrs...)
	return coloring.PrintfFunc()
}

// NoColoredPrinterFunc creates and returns a ColoredPrinterFunc that prints text without any
// color attributes.
func NoColoredPrinterFunc() ColoredPrinterFunc {
	return func(format string, a ...any) {
		fmt.Printf(format, a...)
	}
}
