package cmd

import (
	stdlog "log"
	"os"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cmd/generate"
	"github.com/cloudquery/cloudquery/logging"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Injected with at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"

var (
	Version = "development"
	Commit  = "development"
	Date    = time.Now().String()

	loggerConfig logging.Config

	rootCmd = &cobra.Command{
		Use:     "cloudquery",
		Short:   "cloudquery CLI",
		Version: Version,
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
	rootCmd.AddCommand(generate.Cmd)
	rootCmd.PersistentFlags().BoolVarP(&loggerConfig.Verbose, "verbose", "v", false, "Enable Verbose logging")
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.ConsoleLoggingEnabled, "enableConsoleLog", true, "Enable console logging")
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.EncodeLogsAsJson, "encodeLogsAsJson", false, "EncodeLogsAsJson makes the logging framework logging JSON")
	rootCmd.PersistentFlags().BoolVar(&loggerConfig.FileLoggingEnabled, "enableFileLogging", true, "enableFileLogging makes the framework logging to a file")
	rootCmd.PersistentFlags().StringVar(&loggerConfig.Directory, "logDirectory", ".", "Directory to logging to to when file logging is enabled")
	rootCmd.PersistentFlags().StringVar(&loggerConfig.Filename, "logFile", "cloudquery.log", "Filename is the name of the logfile which will be placed inside the directory")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxSize, "maxSize", 30, "MaxSize the max size in MB of the logfile before it's rolled")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxBackups, "maxBackups", 3, "MaxBackups the max number of rolled files to keep")
	rootCmd.PersistentFlags().IntVar(&loggerConfig.MaxAge, "maxAge", 3, "MaxAge the max age in days to keep a logfile")
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "."
	}
	rootCmd.PersistentFlags().String( "plugin-dir", workingDir, "Directory to save and load Cloudquery plugins from (env: CQ_PLUGIN_DIR)")
	_ = viper.BindPFlag("plugin-dir", rootCmd.PersistentFlags().Lookup("plugin-dir"))
	cobra.OnInitialize(initConfig, initLogging)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func initLogging() {
	zerolog.Logger = logging.Configure(loggerConfig)
}
