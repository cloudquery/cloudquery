package cmd

import (
	stdlog "log"
	"os"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client"

	"github.com/thoas/go-funk"

	"github.com/cloudquery/cloudquery/internal/logging"

	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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

// Injected with at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"

var (
	Commit = "development"
	Date   = time.Now().String()

	loggerConfig logging.Config

	rootCmd = &cobra.Command{
		Use:   "cloudquery",
		Short: "CloudQuery CLI",
		Long: `CloudQuery CLI

Query your cloud assets & configuration with SQL for monitoring security, compliance & cost purposes.

Find more information at:
	https://docs.cloudquery.io`,
		Version: client.Version,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		stdlog.Println(err)
		os.Exit(1)
	}
}

func init() {
	// add inner commands
	rootCmd.PersistentFlags().String("config", "./config.hcl", "path to configuration file. can be generated with 'init {provider}' command (env: CQ_CONFIG_PATH)")
	rootCmd.PersistentFlags().Bool("no-verify", false, "NoVerify is true registry won't verify the plugins")
	rootCmd.PersistentFlags().String("dsn", "", "database connection string (env: CQ_DSN) (example: 'postgres://postgres:pass@localhost:5432/postgres')")
	// Logging Flags
	rootCmd.PersistentFlags().BoolVarP(&loggerConfig.Verbose, "verbose", "v", false, "Enable Verbose logging")
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.ConsoleLoggingEnabled, "enable-console-log", false, "Enable console logging")
	_ = viper.BindPFlag("enable-console-log", rootCmd.PersistentFlags().Lookup("enable-console-log"))
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.EncodeLogsAsJson, "encode-json", false, "EncodeLogsAsJson makes the logging framework logging JSON instead of KV")
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.FileLoggingEnabled, "enable-file-logging", true, "enableFileLogging makes the framework logging to a file")
	rootCmd.PersistentFlags().StringVar(&loggerConfig.Directory, "log-directory", ".", "Directory to logging to to when file logging is enabled")
	rootCmd.PersistentFlags().StringVar(&loggerConfig.Filename, "log-file", "cloudquery.log", "Filename is the name of the logfile which will be placed inside the directory")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxSize, "max-size", 30, "MaxSize the max size in MB of the logfile before it's rolled")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxBackups, "max-backups", 3, "MaxBackups the max number of rolled files to keep")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxAge, "max-age", 3, "MaxAge the max age in days to keep a logfile")
	rootCmd.PersistentFlags().String("plugin-dir", "./.cq/providers", "Directory to save and load CloudQuery plugins from (env: CQ_PLUGIN_DIR)")
	rootCmd.PersistentFlags().String("policy-dir", "./.cq/policies", "Directory to save and load CloudQuery policies from (env: CQ_POLICY_DIR)")
	rootCmd.PersistentFlags().String("reattach-providers", "", "Path to reattach unmanaged plugins, mostly used for testing purposes (env: CQ_REATTACH_PROVIDERS)")
	rootCmd.PersistentFlags().Bool("skip-build-tables", false, "Skip building tables on run, this should only be true if tables already exist.")
	rootCmd.PersistentFlags().Bool("no-telemetry", false, "NoTelemetry is true telemetry collection will be disabled")
	rootCmd.PersistentFlags().Bool("inspect-telemetry", false, "Enable telemetry inspection")
	rootCmd.PersistentFlags().Bool("debug-telemetry", false, "DebugTelemetry is true to debug telemetry logging")
	rootCmd.PersistentFlags().String("telemetry-endpoint", "telemetry.cloudquery.io:443", "Telemetry endpoint")
	rootCmd.PersistentFlags().Bool("insecure-telemetry-endpoint", false, "Allow insecure connection to telemetry endpoint")

	_ = rootCmd.PersistentFlags().MarkHidden("telemetry-endpoint")
	_ = rootCmd.PersistentFlags().MarkHidden("insecure-telemetry-endpoint")

	_ = viper.BindPFlag("plugin-dir", rootCmd.PersistentFlags().Lookup("plugin-dir"))
	_ = viper.BindPFlag("policy-dir", rootCmd.PersistentFlags().Lookup("policy-dir"))
	_ = viper.BindPFlag("reattach-providers", rootCmd.PersistentFlags().Lookup("reattach-providers"))
	_ = viper.BindPFlag("dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	_ = viper.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("no-verify", rootCmd.PersistentFlags().Lookup("no-verify"))
	_ = viper.BindPFlag("skip-build-tables", rootCmd.PersistentFlags().Lookup("skip-build-tables"))
	_ = viper.BindPFlag("no-telemetry", rootCmd.PersistentFlags().Lookup("no-telemetry"))
	_ = viper.BindPFlag("inspect-telemetry", rootCmd.PersistentFlags().Lookup("inspect-telemetry"))
	_ = viper.BindPFlag("debug-telemetry", rootCmd.PersistentFlags().Lookup("debug-telemetry"))
	_ = viper.BindPFlag("telemetry-endpoint", rootCmd.PersistentFlags().Lookup("telemetry-endpoint"))
	_ = viper.BindPFlag("insecure-telemetry-endpoint", rootCmd.PersistentFlags().Lookup("insecure-telemetry-endpoint"))

	registerSentryFlags(rootCmd)

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SetUsageTemplate(usageTemplate)
	cobra.OnInitialize(initConfig, initLogging, initUlimit, initSentry)
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
	zerolog.Logger = logging.Configure(loggerConfig)
}
