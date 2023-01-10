package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

const sentryDsnDefault = "https://3d2f1b94bdb64884ab1a52f56ce56652@o1396617.ingest.sentry.io/6720193"

var (
	Version   = "development"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration at scale.

Find more information at:
	https://www.cloudquery.io`

	disableSentry   = false
	analyticsClient *AnalyticsClient
	logFile         *os.File
)

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	telemetryLevel := enum.NewEnum([]string{"none", "errors", "stats", "all"}, "all")
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"
	sentryDsn := sentryDsnDefault

	// support legacy telemetry environment variable,
	// but the newer CQ_TELEMETRY_LEVEL environment variable takes precedence
	defaultTelemetryValue := telemetryLevel.Value
	legacyTelemetry := os.Getenv("CQ_NO_TELEMETRY")
	if legacyTelemetry != "" {
		defaultTelemetryValue = "none"
	}
	err := telemetryLevel.Set(getEnvOrDefault("CQ_TELEMETRY_LEVEL", defaultTelemetryValue))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to set telemetry level: "+err.Error())
		os.Exit(1)
	}

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
			var err error
			if logFile, err = initLogging(noLogFile, logLevel, logFormat, logConsole, logFileName); err != nil {
				return err
			}

			// log warnings now that the logger is initialized
			if legacyTelemetry != "" {
				log.Warn().Msg("The CQ_NO_TELEMETRY environment variable will be deprecated, please use CQ_TELEMETRY_LEVEL=none instead.")
			}

			sendStats := funk.ContainsString([]string{"all", "stats"}, telemetryLevel.String())
			if Version != "development" && sendStats {
				analyticsClient, err = initAnalytics()
				if err != nil {
					log.Warn().Err(err).Msg("failed to initialize analytics client")
				}
			}

			sendErrors := funk.ContainsString([]string{"all", "errors"}, telemetryLevel.String())
			if sentryDsn != "" && Version != "development" && sendErrors {
				if err := initSentry(sentryDsn, Version); err != nil {
					// we don't fail on sentry init errors as there might be no connection or sentry can be blocked.
					log.Warn().Err(err).Msg("failed to initialize sentry")
				}
			} else {
				disableSentry = true
			}

			return nil
		},
	}

	cmd.PersistentFlags().String("cq-dir", ".cq", "directory to store cloudquery files, such as downloaded plugins")
	cmd.PersistentFlags().String("data-dir", "", "set persistent data directory")
	err = cmd.PersistentFlags().MarkDeprecated("data-dir", "use cq-dir instead")
	if err != nil {
		panic(err)
	}

	cmd.PersistentFlags().String("color", "auto", "Enable colorized output when log-console is set (on, off, auto)")
	err = cmd.PersistentFlags().MarkDeprecated("color", "console logs are always colorless")
	if err != nil {
		panic(err)
	}

	// Logging Flags
	cmd.PersistentFlags().BoolVar(&logConsole, "log-console", false, "enable console logging")
	cmd.PersistentFlags().Var(logFormat, "log-format", "Logging format (json, text)")
	cmd.PersistentFlags().Var(logLevel, "log-level", "Logging level")
	cmd.PersistentFlags().BoolVar(&noLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&logFileName, "log-file-name", "cloudquery.log", "Log filename")

	// Telemetry (analytics) flags
	f := cmd.PersistentFlags().VarPF(telemetryLevel, "telemetry-level", "", "Telemetry level (none, errors, stats, all)")
	f.DefValue = "all"

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(
		NewCmdSync(),
		NewCmdMigrate(),
		newCmdDoc(),
	)
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true
	cobra.OnFinalize(func() {
		if analyticsClient != nil {
			analyticsClient.Close()
		}
	})

	return cmd
}

// formats a timestamp in UTC and RFC3339
func formatTimestampUtcRfc3339(timestamp any) string {
	timestampConcrete, ok := timestamp.(time.Time)
	if !ok {
		return fmt.Sprintf("%v", timestamp)
	}

	return timestampConcrete.UTC().Format(time.RFC3339)
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
