package cmd

import (
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cmd/fetch"
	"github.com/cloudquery/cloudquery/cmd/flags"
	"github.com/cloudquery/cloudquery/cmd/generate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Values for Commit and Date should be injected at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"

	Commit    = "development"
	Date      = "unknown"
	APIKey    = ""
	Version   = "dev"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Open source data integration platform for infrastructure teams.

Find more information at:
	https://cloudquery.io`
)

func newCmdRoot() *cobra.Command {
	// logLevel := newEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	// logFormat := newEnum([]string{"text", "json"}, "text")
	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true

			zerologLevel, err := zerolog.ParseLevel("debug")
			if err != nil {
				return err
			}
			// var logger zerolog.Logger
			// if viper.Get(flags.LogFormat) == "json" {
			log.Logger = zerolog.New(os.Stderr).Level(zerologLevel)
			// } else {
			// 	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerologLevel)
			// }

			// log.Logger = logger
			// log.Logger

			// if !viper.GetBool(flags.NoTelemetry) {
			// 	fmt.Println("Anonymous telemetry collection and crash reporting enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/cli/telemetry")
			// }
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// analytics.Close()
		},
	}

	cmd.PersistentFlags().String(flags.DataDir, "./.cq", "set persistent data directory (env: CQ_DATA_DIR)")
	cmd.PersistentFlags().String(flags.Color, "auto", "Enable colorized output (on, off, auto)")

	// Logging Flags
	cmd.PersistentFlags().BoolP(flags.Verbose, "v", false, "enable verbose logging")
	cmd.PersistentFlags().Bool(flags.LogConsole, false, "enable console logging")
	cmd.PersistentFlags().String(flags.LogFormat, "text", "Logging format (json, text)")
	cmd.PersistentFlags().Bool(flags.NoLogFile, false, "Disable logging to file")
	cmd.PersistentFlags().String(flags.LogFileName, "cloudquery.log", "Log filename")

	// Telemtry (analytics) flags
	cmd.PersistentFlags().Bool(flags.NoTelemetry, false, "disable telemetry collection")
	cmd.PersistentFlags().Bool(flags.TelemetryInspect, false, "enable telemetry inspection")
	cmd.PersistentFlags().Bool(flags.TelemtryDebug, false, "enable telemetry debug logging")

	// Sentry (error reporting) flags
	cmd.PersistentFlags().Bool(flags.SentryDebug, false, "enable Sentry debug mode")
	cmd.PersistentFlags().String(flags.SentryDSN, "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	hiddenFlags := []string{
		flags.TelemetryInspect, flags.TelemtryDebug,
		flags.SentryDebug, flags.SentryDSN}
	for _, f := range hiddenFlags {
		err := cmd.PersistentFlags().MarkHidden(f)
		if err != nil {
			panic(err)
		}
	}
	initViper()
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(generate.NewCmdInit(), fetch.NewCmdFetch())
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
