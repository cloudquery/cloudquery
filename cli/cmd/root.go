package cmd

import (
	"os"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const sentryDsnDefault = "https://3d2f1b94bdb64884ab1a52f56ce56652@o1396617.ingest.sentry.io/6720193"

var (
	Version   = "development"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration at scale.

Find more information at:
	https://cloudquery.io`

	analyticsClient *AnalyticsClient
)

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	telemetryLevel := enum.NewEnum([]string{"none", "errors", "all"}, "all")
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"
	sentryDsn := sentryDsnDefault

	telemetryLevel.Set(getEnvOrDefault("CQ_TELEMETRY_LEVEL", telemetryLevel.Value))

	var logFile *os.File
	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true
			var err error
			if logFile, err = initLogging(noLogFile, logLevel, logFormat, logConsole, logFileName); err != nil {
				return err
			}

			if telemetryLevel.String() == "all" && Version != "development" {
				analyticsClient, err = initAnalytics()
				if err != nil {
					log.Warn().Err(err).Msg("failed to initialize analytics client")
				}
			}

			if sentryDsn != "" && Version != "development" && telemetryLevel.String() != "none" {
				if err := initSentry(sentryDsn, Version); err != nil {
					// we don't fail on sentry init errors as there might be no connection or sentry can be blocked.
					log.Warn().Err(err).Msg("failed to initialize sentry")
				}
			}

			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if logFile != nil {
				logFile.Close()
			}
			if analyticsClient != nil {
				analyticsClient.Close()
			}
		},
	}

	cmd.PersistentFlags().String("data-dir", "./.cq", "set persistent data directory (env: CQ_DATA_DIR)")

	cmd.PersistentFlags().String("color", "auto", "Enable colorized output (on, off, auto)")

	// Logging Flags
	cmd.PersistentFlags().BoolVar(&logConsole, "log-console", false, "enable console logging")
	cmd.PersistentFlags().Var(logFormat, "log-format", "Logging format (json, text)")
	cmd.PersistentFlags().Var(logLevel, "log-level", "Logging level")
	cmd.PersistentFlags().BoolVar(&noLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&logFileName, "log-file-name", "cloudquery.log", "Log filename")

	// Telemetry (analytics) flags
	cmd.PersistentFlags().Var(telemetryLevel, "telemetry-level", "Telemetry level (none, errors, all)")

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(NewCmdSync(), newCmdDoc())
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true
	return cmd
}
