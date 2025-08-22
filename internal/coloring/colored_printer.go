package coloring

import (
	"fmt"

	"github.com/fatih/color"
)

type ColoredPrinterFunc func(format string, a ...any)

func NewColoredPrinterFunc(attrs ...color.Attribute) ColoredPrinterFunc {
	coloring := color.New(attrs...)
	return coloring.PrintfFunc()
}

func NoColoredPrinterFunc() ColoredPrinterFunc {
	return func(format string, a ...any) {
		fmt.Printf(format, a...)
	}
}
