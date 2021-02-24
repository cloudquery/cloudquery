package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
)

// Configuration for logging
type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool
	// Enable Verbose logging
	Verbose bool
	// EncodeLogsAsJson makes the logging framework logging JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework logging to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to logging to to when file logging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
	// Console logging will be without color, console logging must be enabled first.
	ConsoleNoColor bool
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
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
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

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
			Bool("verbose", config.Verbose).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	return logger
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
