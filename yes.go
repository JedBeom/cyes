package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var msg string
	if len(os.Args) == 1 {
		msg = "y\n"
	} else {
		msg = strings.Join(os.Args[1:], " ")
		msg += "\n"
	}
	print := os.Stdout

	red := color.New(color.FgRed).FprintFunc()
	blue := color.New(color.FgBlue).FprintFunc()
	green := color.New(color.FgGreen).FprintFunc()
	yellow := color.New(color.FgYellow).FprintFunc()
	magenta := color.New(color.FgMagenta).FprintFunc()
	cyan := color.New(color.FgCyan).FprintFunc()
    white := color.New(color.FgWhite).FprintFunc()

	for {
		red(print, msg)
		blue(print, msg)
		green(print, msg)
		yellow(print, msg)
		magenta(print, msg)
		cyan(print, msg)
	}
}
