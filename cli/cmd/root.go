package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/rs/zerolog/log"
	"github.com/rudderlabs/analytics-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const sentryDsnDefault = "https://3d2f1b94bdb64884ab1a52f56ce56652@o1396617.ingest.sentry.io/6720193"

var (
	Version   = "development"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration that works.

Find more information at:
	https://cloudquery.io`
	telemetryLevels = []string{"none", "errors", "all"}

	analyticsClient *Analytics
)

func strInArray(str string, arr []string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"
	sentryDsn := sentryDsnDefault

	var logFile *os.File
	var analyticsClient analytics.Client
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

			telemetryLevel := viper.GetString("telemetry-level")
			if !strInArray(telemetryLevel, telemetryLevels) {
				return fmt.Errorf("invalid telemetry level %s. must be one of %v", telemetryLevel, telemetryLevels)
			}
			if telemetryLevel == "all" {
				analyticsClient = initAnalytics()
			}

			if sentryDsn != "" && Version != "development" && telemetryLevel != "none" {
				if err := initSentry(sentryDsn, Version); err != nil {
					// we don't fail on sentry init errors as there might be no connection or sentry can be blocked.
					log.Err(err).Msg("failed to initialize sentry")
				}
			}

			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if logFile != nil {
				logFile.Close()
			}
			analyticsClient.Close()
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

	// Telemtry (analytics) flags
	cmd.PersistentFlags().String("telemetry-level", "all", "Telemetry level (none, errors, all)")
	// we dont need viper support for most flags as all can be used via command line for now (we can add in the future if really necessary)
	// the only exception is the telemetry as people might want to put in a bash starter script
	if err := viper.BindPFlag("telemetry-level", cmd.PersistentFlags().Lookup("telemetry-level")); err != nil {
		panic(err)
	}

	initViper()
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(NewCmdSync(), newCmdDoc())
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true
	return cmd
}

func initViper() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/cq/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.cq") // call multiple times to add many search paths
	viper.AddConfigPath(".")
}
