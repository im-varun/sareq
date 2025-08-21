package coloring

import (
	"fmt"

	"github.com/fatih/color"
)

type ColoredPrinterFunc func(format string, a ...any)

func NewColoredPrinterFunc(attrs ...color.Attribute) ColoredPrinterFunc {
	c := color.New(attrs...)
	return c.PrintfFunc()
}

func NoColoredPrinterFunc() ColoredPrinterFunc {
	return func(format string, a ...any) {
		fmt.Printf(format, a...)
	}
}
