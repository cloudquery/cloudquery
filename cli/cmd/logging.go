package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/cloudquery/cloudquery/cli/v6/internal/enum"
	"github.com/cloudquery/cloudquery/cli/v6/internal/secrets"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogging(noLogFile bool, logLevel *enum.Enum, logFormat *enum.Enum, logConsole bool, logFileName string) (*os.File, error) {
	var logFile *os.File
	zerologLevel, err := zerolog.ParseLevel(logLevel.String())
	if err != nil {
		return nil, err
	}
	var writers []io.Writer
	if !noLogFile {
		logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		if logFormat.String() == "text" {
			// for file logging we don't need color. we can add it as an option but don't think it is useful
			writers = append(writers, zerolog.ConsoleWriter{
				Out:             logFile,
				NoColor:         true,
				FormatTimestamp: formatTimestampUtcRfc3339,
			})
		} else {
			writers = append(writers, logFile)
		}
	}
	if logConsole {
		if err := os.Stdout.Close(); err != nil {
			return nil, fmt.Errorf("failed to close stdout: %w", err)
		}
		if logFormat.String() == "text" {
			writers = append(writers, zerolog.ConsoleWriter{
				Out:             os.Stderr,
				NoColor:         true,
				FormatTimestamp: formatTimestampUtcRfc3339,
			})
		} else {
			writers = append(writers, os.Stderr)
		}
	}
	mw := io.MultiWriter(writers...)
	secretAwareWriter := secrets.NewSecretAwareWriter(mw, secretAwareRedactor)
	log.Logger = zerolog.New(secretAwareWriter).Level(zerologLevel).With().Str("module", "cli").Str("invocation_id", invocationUUID.String()).Timestamp().Logger()
	return logFile, nil
}
