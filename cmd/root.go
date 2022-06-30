package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cmd/completion"
	"github.com/cloudquery/cloudquery/cmd/fetch"
	initCmd "github.com/cloudquery/cloudquery/cmd/init"
	"github.com/cloudquery/cloudquery/cmd/options"
	"github.com/cloudquery/cloudquery/cmd/policy"
	"github.com/cloudquery/cloudquery/cmd/provider"
	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/getsentry/sentry-go"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// fileDescriptorF gets set trough system relevant files like ulimit_unix.go
var fileDescriptorF func()

const (
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
	rootCmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: core.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true

			if analytics.Enabled() {
				ui.ColorizedOutput(ui.ColorInfo, "Anonymous telemetry collection and crash reporting enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/cli/telemetry\n")
				if ui.IsTerminal() {
					if err := helpers.Sleep(cmd.Context(), 2*time.Second); err != nil {
						return err
					}
				}
			}
			logInvocationParams(cmd, args)
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			analytics.Close()
		},
	}

	// add inner commands
	rootCmd.PersistentFlags().String("config", "./cloudquery.yml", "path to configuration file. can be generated with 'init {provider}' command (env: CQ_CONFIG_PATH)")
	rootCmd.PersistentFlags().Bool("no-verify", false, "disable plugins verification")
	rootCmd.PersistentFlags().String("dsn", "", "database connection string (env: CQ_DSN) (example: 'postgres://postgres:pass@localhost:5432/postgres')")

	// Logging Flags
	rootCmd.PersistentFlags().BoolVarP(&logging.GlobalConfig.Verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&logging.GlobalConfig.ConsoleLoggingEnabled, "enable-console-log", false, "enable console logging")
	rootCmd.PersistentFlags().BoolVar(&logging.GlobalConfig.EncodeLogsAsJson, "encode-json", false, "enable JSON log format, instead of key/value")
	rootCmd.PersistentFlags().BoolVar(&logging.GlobalConfig.FileLoggingEnabled, "enable-file-logging", true, "enable file logging")
	rootCmd.PersistentFlags().BoolVar(&logging.GlobalConfig.ConsoleNoColor, "disable-log-color", false, "disable log colors")
	rootCmd.PersistentFlags().StringVar(&logging.GlobalConfig.Directory, "log-directory", ".", "set output directory for logs")
	rootCmd.PersistentFlags().StringVar(&logging.GlobalConfig.Filename, "log-file", "cloudquery.log", "set output filename for logs")
	rootCmd.PersistentFlags().IntVar(&logging.GlobalConfig.MaxSize, "max-size", 30, "set max size in MB of the logfile before it's rolled")
	rootCmd.PersistentFlags().IntVar(&logging.GlobalConfig.MaxBackups, "max-backups", 3, "set max number of rolled files to keep")
	rootCmd.PersistentFlags().IntVar(&logging.GlobalConfig.MaxAge, "max-age", 3, "set max age in days to keep a logfile")
	rootCmd.PersistentFlags().String("data-dir", "./.cq", "set persistent data directory (env: CQ_DATA_DIR)")
	rootCmd.PersistentFlags().String("reattach-providers", "", "path to reattach unmanaged plugins, mostly used for testing purposes (env: CQ_REATTACH_PROVIDERS)")
	rootCmd.PersistentFlags().Bool("skip-build-tables", false, "enable skipping building tables. Should only be set if tables already exist")
	rootCmd.PersistentFlags().Bool("force-drop", false, "when upgrading schema, force dropping of any dependent views")

	rootCmd.PersistentFlags().Bool("no-telemetry", false, "disable telemetry collection")
	rootCmd.PersistentFlags().Bool("inspect-telemetry", false, "enable telemetry inspection")
	rootCmd.PersistentFlags().Bool("debug-telemetry", false, "enable telemetry debug logging")
	rootCmd.PersistentFlags().String("telemetry-apikey", APIKey, "set telemetry API Key")

	_ = rootCmd.PersistentFlags().MarkHidden("inspect-telemetry")
	_ = rootCmd.PersistentFlags().MarkHidden("debug-telemetry")
	_ = rootCmd.PersistentFlags().MarkHidden("telemetry-apikey")

	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag("enable-console-log", rootCmd.PersistentFlags().Lookup("enable-console-log"))
	_ = viper.BindPFlag("data-dir", rootCmd.PersistentFlags().Lookup("data-dir"))
	_ = viper.BindPFlag("reattach-providers", rootCmd.PersistentFlags().Lookup("reattach-providers"))
	_ = viper.BindPFlag("dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	_ = viper.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("no-verify", rootCmd.PersistentFlags().Lookup("no-verify"))
	_ = viper.BindPFlag("skip-build-tables", rootCmd.PersistentFlags().Lookup("skip-build-tables"))
	_ = viper.BindPFlag("force-drop", rootCmd.PersistentFlags().Lookup("force-drop"))

	// Telemetry specific options
	_ = viper.BindPFlag("no-telemetry", rootCmd.PersistentFlags().Lookup("no-telemetry"))
	_ = viper.BindPFlag("debug-telemetry", rootCmd.PersistentFlags().Lookup("debug-telemetry"))
	_ = viper.BindPFlag("inspect-telemetry", rootCmd.PersistentFlags().Lookup("inspect-telemetry"))
	_ = viper.BindPFlag("telemetry-apikey", rootCmd.PersistentFlags().Lookup("telemetry-apikey"))

	registerSentryFlags(rootCmd)

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cobra.OnInitialize(initConfig, initLogging, initUlimit, initSentry, initAnalytics)
	rootCmd.AddCommand(
		initCmd.NewCmdInit(), fetch.NewCmdFetch(), policy.NewCmdPolicy(), provider.NewCmdProvider(),
		options.NewCmdOptions(), completion.NewCmdCompletion(), newCmdVersion(), newCmdDoc())
	return rootCmd
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

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func initLogging() {
	if funk.ContainsString(os.Args, "completion") {
		return
	}
	if !ui.IsTerminal() {
		logging.GlobalConfig.ConsoleLoggingEnabled = true // always true when no terminal
	}
	logging.GlobalConfig.InstanceId = utils.InstanceId.String()

	zerolog.Logger = logging.Configure(logging.GlobalConfig).With().Logger()
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
