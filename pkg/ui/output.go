package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// ColorizedOutput outputs a colored message directly to the terminal.
// The remaining arguments should be interpolations for the format string.
func ColorizedOutput(c *color.Color, msg string, values ...interface{}) {
	if !IsTerminal() {
		// Print output to log
		log.Info().Msgf(strings.ReplaceAll(msg, "\n", ""), values...)
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

func Colorize(c *color.Color, noColor bool, msg string, values ...interface{}) string {
	if noColor {
		return fmt.Sprintf(msg, values...)
	}
	return c.Sprintf(msg, values...)
}

var (
	ColorTrace        = color.New(color.FgMagenta, color.Bold)
	ColorDebug        = color.New(color.FgWhite, color.Faint)
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
	ColorUnderline    = color.New(color.Underline)
)
