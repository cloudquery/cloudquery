package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GlobalConfig is the global alterable logging config
var GlobalConfig Config

// Config for logging
type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool `yaml:"enable_console_logging,omitempty" json:"enable_console_logging,omitempty"`
	// Enable Verbose logging
	Verbose bool `yaml:"verbose,omitempty" json:"verbose,omitempty"`
	// EncodeLogsAsJson makes the logging framework logging JSON
	EncodeLogsAsJson bool `yaml:"encode_logs_as_json,omitempty" json:"encode_logs_as_json,omitempty"`
	// FileLoggingEnabled makes the framework logging to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool `yaml:"file_logging_enabled,omitempty" json:"file_logging_enabled,omitempty"`
	// Directory to logging to to when file logging is enabled
	Directory string `yaml:"directory,omitempty" json:"directory,omitempty"`
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string `yaml:"filename,omitempty" json:"filename,omitempty"`
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int `yaml:"max_size,omitempty" json:"max_size,omitempty"`
	// MaxBackups the max number of rolled files to keep
	MaxBackups int `yaml:"max_backups,omitempty" json:"max_backups,omitempty"`
	// MaxAge the max age in days to keep a logfile
	MaxAge int `yaml:"max_age,omitempty" json:"max_age,omitempty"`
	// Console logging will be without color, console logging must be enabled first.
	ConsoleNoColor bool `yaml:"console_no_color,omitempty" json:"console_no_color,omitempty"`
	// Unique Identifier of execution
	InstanceId string

	// console is a writer that will be used for console output. If it is not set os.Stderr will be used.
	console io.Writer
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output logging file should be located at /var/logging/service-xyz/service-xyz.logging and
// will be rolled according to configuration set.
func Configure(config Config) zerolog.Logger {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		if config.EncodeLogsAsJson {
			writers = append(writers, os.Stdout)
		} else {
			console := config.console
			if console == nil {
				console = os.Stderr
			}
			writers = append(writers, zerolog.ConsoleWriter{FormatLevel: formatLevel(config.ConsoleNoColor), Out: console, NoColor: config.ConsoleNoColor})
		}
	}

	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	// Default level is info, unless verbose flag is on
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(mw).With().Timestamp().Str("instance_id", config.InstanceId).Logger()
	// override global logger
	log.Logger = logger
	// Default level is info, unless verbose flag is on
	logger.Level(zerolog.InfoLevel)
	if config.Verbose {
		logger.Level(zerolog.DebugLevel)
	}

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Bool("consoleLog", config.ConsoleLoggingEnabled).
		Bool("verbose", config.Verbose).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	return logger
}

// Reconfigure reconfigures the already initialized Logger with new values
func Reconfigure(originalConfig, updatedConfig Config) {
	// FIXME these look buggy
	if updatedConfig.Verbose {
		originalConfig.Verbose = updatedConfig.Verbose
	}
	if updatedConfig.ConsoleLoggingEnabled {
		originalConfig.ConsoleLoggingEnabled = updatedConfig.ConsoleLoggingEnabled
	}
	if updatedConfig.EncodeLogsAsJson {
		originalConfig.EncodeLogsAsJson = updatedConfig.EncodeLogsAsJson
	}
	if !updatedConfig.FileLoggingEnabled {
		originalConfig.FileLoggingEnabled = updatedConfig.FileLoggingEnabled
	}
	if updatedConfig.ConsoleNoColor {
		originalConfig.ConsoleNoColor = updatedConfig.ConsoleNoColor
	}
	log.Logger = Configure(GlobalConfig)
}

func newRollingFile(config Config) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create logging directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
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
