package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const sentryDsnDefault = "https://3d2f1b94bdb64884ab1a52f56ce56652@o1396617.ingest.sentry.io/6720193"

var (
	Version   = "development"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration that works.

Find more information at:
	https://cloudquery.io`
)

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"
	sentryDsn := sentryDsnDefault
	noTelemetry := false

	var logFile *os.File
	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			zerolog.TimestampFunc = func() time.Time {
				return time.Now().UTC()
			}

			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true
			zerologLevel, err := zerolog.ParseLevel(logLevel.String())
			if err != nil {
				return err
			}
			var writers []io.Writer
			if !noLogFile {
				logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					return err
				}
				if logFormat.String() == "text" {
					// for file logging we dont need color. we can add it as an option but don't think it is useful
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
					return fmt.Errorf("failed to close stdout: %w", err)
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
			log.Logger = zerolog.New(mw).Level(zerologLevel).With().Str("module", "cli").Timestamp().Logger()

			noTelemetry = getEnvOrDefault("CQ_NO_TELEMETRY", "false") == "true" || noTelemetry
			sentryEnabled := sentryDsn != "" && Version != "development" && !noTelemetry
			if sentryEnabled {
				if err := sentry.Init(sentry.ClientOptions{
					Debug:     false,
					Dsn:       sentryDsn,
					Release:   "cloudquery@" + Version,
					Transport: sentry.NewHTTPSyncTransport(),
					// https://docs.sentry.io/platforms/go/configuration/options/#removing-default-integrations
					Integrations: func(integrations []sentry.Integration) []sentry.Integration {
						var filteredIntegrations []sentry.Integration
						for _, integration := range integrations {
							if integration.Name() == "Modules" {
								continue
							}
							filteredIntegrations = append(filteredIntegrations, integration)
						}
						return filteredIntegrations
					},
				}); err != nil {
					return err
				}
			}

			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if logFile != nil {
				logFile.Close()
			}
		},
	}

	cmd.PersistentFlags().String("data-dir", "./.cq", "set persistent data directory (env: CQ_DATA_DIR)")

	cmd.PersistentFlags().String("color", "auto", "Enable colorized output when log-console is set (on, off, auto)")
	err := cmd.PersistentFlags().MarkDeprecated("color", "console logs are always colorless")
	if err != nil {
		panic(err)
	}

	// Logging Flags
	cmd.PersistentFlags().BoolVar(&logConsole, "log-console", false, "enable console logging")
	cmd.PersistentFlags().Var(logFormat, "log-format", "Logging format (json, text)")
	cmd.PersistentFlags().Var(logLevel, "log-level", "Logging level")
	cmd.PersistentFlags().BoolVar(&noLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&logFileName, "log-file-name", "cloudquery.log", "Log filename")

	// Telemtry (analytics) flags
	cmd.PersistentFlags().BoolVar(&noTelemetry, "no-telemetry", false, "disable telemetry collection (env: CQ_NO_TELEMETRY)")
	cmd.PersistentFlags().Bool("telemetry-inspect", false, "enable telemetry inspection")
	cmd.PersistentFlags().Bool("telemetry-debug", false, "enable telemetry debug logging")

	// Sentry (error reporting) flags
	cmd.PersistentFlags().StringVar(&sentryDsn, "sentry-dsn", sentryDsnDefault, "sentry DSN")

	hiddenFlags := []string{"telemetry-inspect", "telemetry-debug", "sentry-dsn"}
	for _, f := range hiddenFlags {
		err := cmd.PersistentFlags().MarkHidden(f)
		if err != nil {
			panic(err)
		}
	}
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(NewCmdSync(), newCmdDoc())
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true
	return cmd
}

// formats a timestamp in UTC and RFC3339
func formatTimestampUtcRfc3339(timestamp interface{}) string {
	timestampConcrete, ok := timestamp.(time.Time)
	if !ok {
		return fmt.Sprintf("%v", timestamp)
	}

	return timestampConcrete.UTC().Format(time.RFC3339)
}
