package ui

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// ColorizedOutput outputs a colored message directly to the terminal.
// The remaining arguments should be interpolations for the format string.
func ColorizedOutput(c *color.Color, msg string, values ...interface{}) {
	if viper.GetBool("enable-console-log") {
		// Print output to log
		if logMsg := strings.ReplaceAll(msg, "\n", ""); logMsg != "" {
			log.Info().Msgf(logMsg, values...)
		}
		return
	}
	_, _ = c.Printf(msg, values...)
}

func IsTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) && term.IsTerminal(int(os.Stdout.Fd()))
}

func DoProgress() bool {
	return IsTerminal() && !viper.GetBool("enable-console-log")
}

func SleepBeforeError(ctx context.Context) {
	if !IsTerminal() {
		return
	}

	select {
	case <-ctx.Done():
	case <-time.After(100 * time.Millisecond):
	}
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
