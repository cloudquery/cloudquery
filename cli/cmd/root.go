package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cli/cmd/enum"
	"github.com/cloudquery/cloudquery/cli/cmd/generate"
	"github.com/cloudquery/cloudquery/cli/cmd/sync"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Values for Commit and Date should be injected at build time with -ldflags "-X github.com/cloudquery/cloudquery/cli/cmd.Variable=Value"

	Commit    = "development"
	Date      = "unknown"
	APIKey    = ""
	Version   = "dev"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration that works.

Find more information at: https://cloudquery.io`
)

func newCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "json")
	noColor := false
	logConsole := false
	noLogFile := false
	logFileName := "cloudquery.log"

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
				logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if logFile != nil {
				logFile.Close()
			}
			// analytics.Close()
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
	viper.BindPFlag("no-telemetry", cmd.PersistentFlags().Lookup("no-telemetry"))
	cmd.PersistentFlags().Bool("telemetry-inspect", false, "enable telemetry inspection")
	cmd.PersistentFlags().Bool("telemetry-debug", false, "enable telemetry debug logging")

	// Sentry (error reporting) flags
	cmd.PersistentFlags().Bool("sentry-debug", false, "enable Sentry debug mode")
	cmd.PersistentFlags().String("sentry-dsn", "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	hiddenFlags := []string{"telemetry-inspect", "telemetry-debug", "sentry-debug", "sentry-dsn"}
	for _, f := range hiddenFlags {
		err := cmd.PersistentFlags().MarkHidden(f)
		if err != nil {
			panic(err)
		}
	}
	initViper()
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(generate.NewCmdGenerate(), sync.NewCmdFetch())
	// cmd.AddCommand(
	// 	initCmd.NewCmdInit(), fetch.NewCmdFetch(), policy.NewCmdPolicy(), provider.NewCmdProvider(),
	// 	options.NewCmdOptions(), newCmdVersion(), newCmdDoc())
	cmd.DisableAutoGenTag = true
	return cmd
}

func Execute() error {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	return newCmdRoot().Execute()
}

func initViper() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}
