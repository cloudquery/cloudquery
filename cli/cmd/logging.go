package cmd

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/cloudquery/cloudquery/cli/v6/internal/otel"
	"io"
	"os"

	"github.com/cloudquery/cloudquery/cli/v6/internal/enum"
	"github.com/cloudquery/cloudquery/cli/v6/internal/secrets"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogging(noLogFile bool, logLevel *enum.Enum, logFormat *enum.Enum, logConsole bool, logFileName string) (*os.File, func(), error) {
	var logFile *os.File
	shutdownFn := func() {}
	zerologLevel, err := zerolog.ParseLevel(logLevel.String())
	if err != nil {
		return nil, shutdownFn, err
	}
	var writers []io.Writer
	if !noLogFile {
		logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, shutdownFn, err
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
			return nil, shutdownFn, fmt.Errorf("failed to close stdout: %w", err)
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

	otelEndpoint := env.GetEnvOrDefault("CLOUD_PLATFORM_OTEL_LOGS_ENDPOINT", "")
	if otelEndpoint != "" {
		shutdownFn, err = otel.SetupOtel(context.Background(), log.Logger, otelEndpoint)
		if err != nil {
			return nil, shutdownFn, fmt.Errorf("failed to setup OpenTelemetry: %w", err)
		}
		log.Logger = log.Logger.Hook(otel.NewOTELLoggerHook())
		log.Info().Str("otel_logs_endpoint", otelEndpoint).Msg("otel logs endpoint set")
	}

	return logFile, shutdownFn, nil
}
