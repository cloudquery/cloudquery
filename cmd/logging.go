package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(options rootOptions) zerolog.Logger {
	var writers []io.Writer

	if options.LogConsole {
		if options.LogFormat.Value == "json" {
			writers = append(writers, os.Stdout)
		} else {
			console := os.Stderr
			writers = append(writers, zerolog.ConsoleWriter{FormatLevel: formatLevel(options.Color.Value == "off"), Out: console, NoColor: options.Color.Value == "off"})
		}
	}

	if !options.NoLogFile {
		writers = append(writers, newRollingFile(options))
	}
	mw := io.MultiWriter(writers...)

	// Default level is info, unless verbose flag is on
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if options.Verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(mw).With().Timestamp().Str("instance_id", util.InstanceId.String()).Logger()
	// override global logger
	log.Logger = logger
	// Default level is info, unless verbose flag is on
	logger.Level(zerolog.InfoLevel)
	if options.Verbose {
		logger.Level(zerolog.DebugLevel)
	}

	logger.Info().
		Bool("noLogFile", options.NoLogFile).
		Str("logFormat", options.LogFormat.Value).
		Bool("logConsole", options.LogConsole).
		Bool("verbose", options.Verbose).
		Str("logDirectory", options.LogDirectory).
		Str("logFilename", options.LogFilename).
		Int("logMaxSize", options.LogMaxSize).
		Int("maxBackups", options.LogMaxBackups).
		Msg("logging configured")

	return logger
}

func newRollingFile(options rootOptions) io.Writer {
	if err := os.MkdirAll(options.LogDirectory, 0744); err != nil {
		log.Error().Err(err).Str("path", options.LogDirectory).Msg("can't create logging directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(options.LogDirectory, options.LogFilename),
		MaxBackups: options.LogMaxBackups, // files
		MaxSize:    options.LogMaxSize,    // megabytes
		MaxAge:     options.LogMaxAge,     // days
	}
}

func formatLevel(noColor bool) func(i interface{}) string {
	// formatLevel is zerolog.Formatter that turns a level value into a string.
	return func(i interface{}) string {
		if level, ok := i.(string); ok {
			switch level {
			case "trace":
				return ui.Colorize(ui.ColorTrace, noColor, "TRC")
			case "debug":
				return ui.Colorize(ui.ColorDebug, noColor, "DBG")
			case "info":
				return ui.Colorize(ui.ColorInfo, noColor, "INF")
			case "warn":
				return ui.Colorize(ui.ColorWarning, noColor, "WRN")
			case "error":
				return ui.Colorize(ui.ColorError, noColor, "ERR")
			case "fatal":
				return ui.Colorize(ui.ColorError, noColor, "FTL")
			case "panic":
				return ui.Colorize(ui.ColorErrorBold, noColor, "PNC")
			default:
				return ui.Colorize(ui.ColorInfo, noColor, "???")
			}
		}
		if i == nil {
			return ui.Colorize(ui.ColorInfo, noColor, "???")
		}
		return strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
	}
}
