package ui

import (
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// ColorizedOutput outputs a colored message directly to the terminal.
// The remaining arguments should be interpolations for the format string.
func ColorizedOutput(c *color.Color, msg string, values ...interface{}) {
	if !IsTerminal() {
		return
	}
	_, _ = c.Printf(msg, values...)
}

func IsTerminal() bool {
	if viper.GetBool("enable-console-log") {
		return false
	}
	return isatty.IsTerminal(os.Stdout.Fd()) && term.IsTerminal(int(os.Stdout.Fd()))
}

var (
	ColorHeader       = color.New(color.Bold)
	ColorInfo         = color.New()
	ColorProgress     = color.New(color.FgCyan)
	ColorProgressBold = color.New(color.FgCyan, color.Bold)
	ColorError        = color.New(color.FgRed)
	ColorErrorBold    = color.New(color.FgRed, color.Bold)
	ColorSuccess      = color.New(color.FgGreen)
	ColorSuccessBold  = color.New(color.FgGreen, color.Bold)
	ColorWarning      = color.New(color.FgYellow)
	ColorWarningBold  = color.New(color.FgYellow, color.Bold)
)
