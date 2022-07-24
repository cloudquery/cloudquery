package cmd

import (
	"io"
	"os"
	"path"

	"github.com/cloudquery/cloudquery/cmd/flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() zerolog.Logger {
	var writers []io.Writer

	if viper.GetBool(flags.LogConsole) {
		if viper.GetString(flags.LogFormat) == "json" {
			writers = append(writers, os.Stdout)
		} else {
			console := os.Stderr
			writers = append(writers,
				zerolog.ConsoleWriter{
					Out:     console,
					NoColor: viper.GetString(flags.Color) == "off"})
		}
	}

	if !viper.GetBool(flags.NoLogFile) {
		writers = append(writers, newRollingFile())
	}
	mw := io.MultiWriter(writers...)

	// Default level is info, unless verbose flag is on
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if viper.GetBool(flags.Verbose) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(mw).With().Timestamp().Str("instance_id", "123").Logger()
	// override global logger
	log.Logger = logger
	// Default level is info, unless verbose flag is on
	logger.Level(zerolog.InfoLevel)
	if viper.GetBool(flags.Verbose) {
		logger.Level(zerolog.DebugLevel)
	}

	logger.Info().
		Bool(flags.NoLogFile, viper.GetBool(flags.NoLogFile)).
		Str(flags.LogFormat, viper.GetString(flags.LogFormat)).
		Bool(flags.LogConsole, viper.GetBool(flags.LogConsole)).
		Bool(flags.Verbose, viper.GetBool(flags.Verbose)).
		Str(flags.LogFileDirectory, viper.GetString(flags.LogFileDirectory)).
		Str(flags.LogFileName, viper.GetString(flags.LogFileName)).
		Int(flags.LogFileMaxAge, viper.GetInt(flags.LogFileMaxAge)).
		Int(flags.LogFileMaxBackups, viper.GetInt(flags.LogFileMaxBackups)).
		Msg("logging configured")

	return logger
}

func newRollingFile() io.Writer {
	if err := os.MkdirAll(viper.GetString(flags.LogFileDirectory), 0744); err != nil {
		log.Error().Err(err).Str("path", viper.GetString(flags.LogFileDirectory)).Msg("can't create logging directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(viper.GetString(flags.LogFileDirectory), viper.GetString(flags.LogFileName)),
		MaxBackups: viper.GetInt(flags.LogFileMaxBackups), // files
		MaxSize:    viper.GetInt(flags.LogFileMaxSize),    // megabytes
		MaxAge:     viper.GetInt(flags.LogFileMaxAge),     // days
	}
}
