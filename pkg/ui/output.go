package ui

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
	"github.com/savioxavier/termlink"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var removeAnsi = regexp.MustCompile(`(?i)\\u00[1-9]b[[0-9;]*[mGKHF]`)
var emojiStatus = []string{color.GreenString("✓"), "📋", color.RedString("❌"), "⚠️", "⌛"}

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

// ColorizedOutput outputs a colored message directly to the terminal.
// The remaining arguments should be interpolations for the format string.
func ColorizedOutput(c *color.Color, msg string, values ...interface{}) {
	if viper.GetBool("enable-console-log") {
		// Print output to log
		if logMsg := strip(fmt.Sprintf(msg, values...)); logMsg != "" {
			switch c {
			case ColorDebug:
				log.Debug().Msg(logMsg)
			case ColorError:
				log.Error().Msg(logMsg)
			case ColorWarning:
				log.Warn().Msg(logMsg)
			case ColorTrace:
				log.Trace().Msg(logMsg)
			default:
				log.Info().Msg(logMsg)
			}
		}
		if IsTerminal() {
			return
		}
	}
	_, _ = c.Printf(msg, values...)
}

func ColorizedNoLogOutput(c *color.Color, msg string, values ...interface{}) {
	if viper.GetBool("enable-console-log") && IsTerminal() {
		return
	}
	// TODO: make zerolog not print this
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

func Link(text string, url string) string {
	if termlink.SupportsHyperlinks() {
		return termlink.Link(text, url)
	}
	// termlink has default behavior for terminals that don't support hyperlinks but it adds an extra space before the link, e.g. `text ( link)`
	// so we use our own formatting
	return fmt.Sprintf("%s (%s)", text, url)
}

func strip(str string) string {
	for _, s := range emojiStatus {
		str = strings.ReplaceAll(str, s, "")
	}
	return strings.TrimRight(strings.TrimSpace(removeAnsi.ReplaceAllString(str, "")), ".")
}
