package console

import (
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
	"os"
)


// ColorizedOutput outputs a colored message directly to the terminal.
// The remaining arguments should be interpolations for the format string.
func ColorizedOutput(c *color.Color, msg string, values ...interface{}) {
	if !IsTerminal() {
		log.WithLevel(colorToLevel(c)).Msgf(msg, values...)
	}
	_, _ = c.Printf(msg, values)
}

func IsTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) && term.IsTerminal(int(os.Stdout.Fd()))
}


func colorToLevel(c *color.Color) zerolog.Level {
	if c.Equals(ColorError) || c.Equals(ColorErrorBold){
		return zerolog.ErrorLevel
	}
	if c.Equals(ColorWarning) || c.Equals(ColorWarningBold){
		return zerolog.WarnLevel
	}
	return zerolog.InfoLevel
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