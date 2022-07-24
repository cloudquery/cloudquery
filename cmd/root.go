package cmd

import (
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cmd/fetch"
	"github.com/cloudquery/cloudquery/cmd/flags"
	initCmd "github.com/cloudquery/cloudquery/cmd/init"
	"github.com/cloudquery/cloudquery/cmd/options"
	"github.com/cloudquery/cloudquery/cmd/policy"
	"github.com/cloudquery/cloudquery/cmd/provider"
	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/getsentry/sentry-go"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// fileDescriptorF gets set trough system relevant files like ulimit_unix.go
var fileDescriptorF func()

var (
	// Values for Commit and Date should be injected at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"

	Commit    = "development"
	Date      = "unknown"
	APIKey    = ""
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

Query your cloud assets & configuration with SQL for monitoring security, compliance & cost purposes.

Find more information at:
	https://docs.cloudquery.io`
)

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: core.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true
			zerolog.Logger = NewLogger()
			initSentry()

			if !viper.GetBool(flags.NoTelemetry) {
				ui.ColorizedOutput(ui.ColorInfo, "Anonymous telemetry collection and crash reporting enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/cli/telemetry\n")
				if ui.IsTerminal() {
					if err := helpers.Sleep(cmd.Context(), 2*time.Second); err != nil {
						return err
					}
				}
				initAnalytics()
			}
			logInvocationParams(cmd, args)
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			analytics.Close()
		},
	}

	cmd.PersistentFlags().String(flags.DataDir, "./.cq", "set persistent data directory (env: CQ_DATA_DIR)")
	cmd.PersistentFlags().String(flags.Color, "auto", "Enable colorized output (on, off, auto)")

	// Logging Flags
	cmd.PersistentFlags().BoolP(flags.Verbose, "v", false, "enable verbose logging")
	cmd.PersistentFlags().Bool(flags.LogConsole, false, "enable console logging")
	cmd.PersistentFlags().String(flags.LogFormat, "keyvalue", "Logging format (json, keyvalue)")
	cmd.PersistentFlags().Bool(flags.NoLogFile, false, "Disable logging to file")
	cmd.PersistentFlags().String(flags.LogFileName, "cloudquery.log", "Log filename")
	cmd.PersistentFlags().String(flags.LogFileDirectory, ".", "Directory to save log files")
	cmd.PersistentFlags().Int(flags.LogFileMaxSize, 30, "Max size in MB of the logfile before it's rolled")
	cmd.PersistentFlags().Int(flags.LogFileMaxBackups, 3, "Max number of rolled files to keep")
	cmd.PersistentFlags().Int(flags.LogFileMaxAge, 3, "Max age in days to keep a logfile")

	// Telemtry (analytics) flags
	cmd.PersistentFlags().Bool(flags.NoTelemetry, false, "disable telemetry collection")
	cmd.PersistentFlags().Bool(flags.TelemetryInspect, false, "enable telemetry inspection")
	cmd.PersistentFlags().Bool(flags.TelemtryDebug, false, "enable telemetry debug logging")

	// Sentry (error reporting) flags
	cmd.PersistentFlags().Bool(flags.SentryDebug, false, "enable Sentry debug mode")
	cmd.PersistentFlags().String(flags.SentryDSN, "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	hiddenFlags := []string{
		flags.TelemetryInspect, flags.TelemtryDebug,
		flags.SentryDebug, flags.SentryDSN,
		flags.LogFileMaxAge, flags.LogFileMaxBackups, flags.LogFileMaxSize}
	for _, f := range hiddenFlags {
		err := cmd.PersistentFlags().MarkHidden(f)
		if err != nil {
			panic(err)
		}
	}
	initViper()
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(
		initCmd.NewCmdInit(), fetch.NewCmdFetch(), policy.NewCmdPolicy(), provider.NewCmdProvider(),
		options.NewCmdOptions(), newCmdVersion(), newCmdDoc())
	cmd.DisableAutoGenTag = true
	return cmd
}

func Execute() error {
	defer func() {
		if err := recover(); err != nil {
			sentry.CurrentHub().Recover(err)
			panic(err)
		}
	}()
	return newCmdRoot().Execute()
}

func initUlimit() {
	if fileDescriptorF != nil {
		fileDescriptorF()
	}
}

func initViper() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func initAnalytics() {
	opts := []analytics.Option{
		analytics.WithVersionInfo(core.Version, Commit, Date),
		analytics.WithTerminal(ui.IsTerminal()),
		analytics.WithApiKey(viper.GetString("telemetry-apikey")),
		analytics.WithInstanceId(utils.InstanceId.String()),
	}
	userId := analytics.GetCookieId()
	if viper.GetBool("no-telemetry") || analytics.CQTeamID == userId.String() {
		opts = append(opts, analytics.WithDisabled())
	}
	if viper.GetBool("debug-telemetry") {
		opts = append(opts, analytics.WithDebug())
	}
	if viper.GetBool("inspect-telemetry") {
		opts = append(opts, analytics.WithInspect())
	}

	_ = analytics.Init(opts...)
}

func logInvocationParams(cmd *cobra.Command, args []string) {
	l := zerolog.Info().Str("core_version", core.Version)
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Name == "dsn" {
			l = l.Str("pflag:"+f.Name, "(redacted)")
			return
		}

		l = l.Str("pflag:"+f.Name, f.Value.String())
	})
	cmd.Flags().Visit(func(f *pflag.Flag) {
		l = l.Str("flag:"+f.Name, f.Value.String())
	})

	l.Str("command", cmd.CommandPath()).Strs("args", args).Msg("Invocation parameters")
}
