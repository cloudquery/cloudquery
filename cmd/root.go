package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cmd/completion"
	"github.com/cloudquery/cloudquery/cmd/fetch"
	initCmd "github.com/cloudquery/cloudquery/cmd/init"
	"github.com/cloudquery/cloudquery/cmd/options"
	"github.com/cloudquery/cloudquery/cmd/policy"
	"github.com/cloudquery/cloudquery/cmd/provider"
	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/cmd/version"
	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/internal/cqpflag"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/helpers"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// This is copied from https://github.com/spf13/cobra/blob/master/command.go#L491
// and modified to not print global flags (as they will be printed via a new options command)
const usageTemplate = `Usage:{{if .Runnable}}
{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
{{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
{{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
Use "{{.CommandPath}} options" for a list of global CLI options.{{end}}
`

// This is copied from https://github.com/spf13/cobra/blob/master/command.go#L491
// and used in the new options command as everywhere else it's disabled via usageTemplate
const usageTemplateWithFlags = `Usage:{{if .Runnable}}
{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
{{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
{{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
Use "{{.Root.Use}} options" for a list of global CLI options.
`

var (
	// Values for Commit and Date should be injected at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"

	instanceId = uuid.New()

	// rootCmd =
)

const (
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

	Query your cloud assets & configuration with SQL for monitoring security, compliance & cost purposes.
	
	Find more information at: https://docs.cloudquery.io`
)

type rootOptions struct {
	Verbose       bool
	Color         *cqpflag.Enum
	Config        string
	LogConsole    bool
	NoLogFile     bool
	LogFormat     *cqpflag.Enum
	LogFilename   string
	LogDirectory  string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
}

func Execute() error {
	defer func() {
		if err := recover(); err != nil {
			sentry.CurrentHub().Recover(err)
			panic(err)
		}
	}()
	return newRootCmd().Execute()
}

func newRootCmd() *cobra.Command {
	o := rootOptions{}
	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: core.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			zerolog.Logger = NewLogger(o)
			initSentry(o)

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
				initAnalytics(o)
			}
			logInvocationParams(cmd, args)
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			analytics.Close()
		},
	}
	cmd.PersistentFlags().String("data-dir", "./.cq", "Directory for providers, policies and other internal CQ data")

	o.Color = cqpflag.NewEnum([]string{"on", "off", "auto"}, "auto")
	cmd.PersistentFlags().Var(o.Color, "color", "Enable colorized output (on, off, auto)")

	// Logging flags
	o.LogFormat = cqpflag.NewEnum([]string{"json", "keyvalue"}, "keyvalue")
	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "Enable verbose logging")
	cmd.PersistentFlags().BoolVar(&o.LogConsole, "log-console", false, "Enable console logging")
	cmd.PersistentFlags().Var(o.LogFormat, "log-format", "Logging format (json, keyvalue)")
	cmd.PersistentFlags().BoolVar(&o.NoLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&o.LogFilename, "log-filename", "cloudquery.log", "Log filename")
	cmd.PersistentFlags().StringVar(&o.LogDirectory, "log-directory", ".", "Directory to save log files")
	cmd.PersistentFlags().IntVar(&o.LogMaxSize, "log-max-size", 30, "Max size in MB of the logfile before it's rolled")
	cmd.PersistentFlags().IntVar(&o.LogMaxBackups, "log-max-backups", 3, "Max number of rolled files to keep")
	cmd.PersistentFlags().IntVar(&o.LogMaxAge, "log-max-age", 3, "Max age in days to keep a logfile")

	// Telemetry (analytics) flags
	cmd.PersistentFlags().Bool("no-telemetry", false, "disable telemetry collection")
	cmd.PersistentFlags().Bool("telemetry-inspect", false, "enable telemetry inspection")
	cmd.PersistentFlags().Bool("telemtry-debug", false, "enable telemetry debug logging")

	// Sentry (error reporting) flags
	cmd.PersistentFlags().Bool("sentry-debug", false, "Enable Sentry debug mode")
	cmd.PersistentFlags().String("sentry-dsn", "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	hiddenFlags := []string{
		"telemetry-inspect", "telemtry-debug",
		"sentry-debug", "sentry-dsn",
		"log-max-age", "log-max-backups", "log-max-size"}
	for _, f := range hiddenFlags {
		err := cmd.PersistentFlags().MarkHidden(f)
		if err != nil {
			panic(err)
		}
	}

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.SetUsageTemplate(usageTemplate)
	cmd.AddCommand(initCmd.NewCmdInit(), fetch.NewCmdFetch(), policy.NewCmdPolicy(), provider.NewCmdProvider(),
		NewCmdDoc(), options.NewCmdOptions(), version.NewCmdVersion(), completion.NewCmdCompletion())
	initConfig()
	bindFlags(cmd, viper.GetViper())
	return cmd
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func initAnalytics(options rootOptions) {
	opts := []analytics.Option{
		analytics.WithVersionInfo(core.Version, util.Commit, util.Date),
		analytics.WithTerminal(ui.IsTerminal()),
		analytics.WithApiKey(viper.GetString("telemetry-apikey")),
		analytics.WithInstanceId(instanceId.String()),
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

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("CQ_%s", envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
