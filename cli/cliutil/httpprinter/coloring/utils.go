package coloring

import "fmt"

type ColoredPrinter func(format string, a ...interface{})

func NoColorPrinter(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
