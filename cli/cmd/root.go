package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
)

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	noColor := false
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"
	sentryDsn := sentryDsnDefault

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
			zerologLevel, err := zerolog.ParseLevel(logLevel.String())
			if err != nil {
				return err
			}
			var writers []io.Writer
			if !noLogFile {
				logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
				if err != nil {
					return err
				}
				if logFormat.String() == "text" {
					// for file logging we dont need color. we can add it as an option but don't think it is useful
					writers = append(writers, zerolog.ConsoleWriter{Out: logFile, NoColor: true})
				} else {
					writers = append(writers, logFile)
				}
			}
			if logConsole {
				if logFormat.String() == "text" {
					writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, NoColor: noColor})
				} else {
					writers = append(writers, os.Stderr)
				}
			}

			mw := io.MultiWriter(writers...)
			log.Logger = zerolog.New(mw).Level(zerologLevel).With().Str("module", "cli").Timestamp().Logger()
			if sentryDsn != "" && Version != "development" {
				if err := sentry.Init(sentry.ClientOptions{
					Debug:   false,
					Dsn:     sentryDsn,
					Release: "cloudquery@" + Version,
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

	cmd.PersistentFlags().String("color", "auto", "Enable colorized output (on, off, auto)")

	// Logging Flags
	cmd.PersistentFlags().BoolVar(&logConsole, "log-console", false, "enable console logging")
	cmd.PersistentFlags().Var(logFormat, "log-format", "Logging format (json, text)")
	cmd.PersistentFlags().Var(logLevel, "log-level", "Logging level")
	cmd.PersistentFlags().BoolVar(&noLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&logFileName, "log-file-name", "cloudquery.log", "Log filename")

	// Telemtry (analytics) flags
	cmd.PersistentFlags().Bool("no-telemetry", false, "disable telemetry collection")
	// we dont need viper support for most flags as all can be used via command line for now (we can add in the future if really necessary)
	// the only exception is the telemetry as people might want to put in a bash starter script
	if err := viper.BindPFlag("no-telemetry", cmd.PersistentFlags().Lookup("no-telemetry")); err != nil {
		panic(err)
	}
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
	initViper()
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(NewCmdGenerate(), NewCmdSync(), newCmdDoc())
	cmd.DisableAutoGenTag = true
	return cmd
}

func initViper() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}
