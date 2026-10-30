package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/enum"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/cloudquery/cloudquery/cli/v6/internal/otel"
	"github.com/cloudquery/cloudquery/cli/v6/internal/secrets"
	"github.com/cloudquery/cloudquery/cli/v6/internal/uuid"
	guuid "github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

const sentryDsnDefault = "https://3d2f1b94bdb64884ab1a52f56ce56652@o1396617.ingest.us.sentry.io/6720193"

var (
	Version   = "development"
	rootShort = "CloudQuery CLI"
	rootLong  = `CloudQuery CLI

High performance data integration at scale.

Find more information at:
	https://www.cloudquery.io`

	disableSentry       = false
	logConsole          = false
	logFile             *os.File
	invocationUUID      uuid.UUID
	secretAwareRedactor *secrets.SecretAwareRedactor
	loggingShutdownFn   func()
)

func NewCmdRoot() *cobra.Command {
	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "text")
	telemetryLevel := enum.NewEnum([]string{"none", "errors", "stats", "all"}, "all")
	noLogFile := false
	logFileName := "cloudquery.log"
	logFileOverwrite := false
	sentryDsn := sentryDsnDefault
	var err error
	if invocationUUID.UUID, err = guuid.NewRandom(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate invocation uuid: %v", err)
		os.Exit(1)
	}
	secretAwareRedactor = secrets.NewSecretAwareRedactor()

	err = telemetryLevel.Set(env.GetEnvOrDefault("CQ_TELEMETRY_LEVEL", telemetryLevel.Value))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to set telemetry level: %v", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:     "cloudquery",
		Short:   rootShort,
		Long:    rootLong,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			zerolog.TimestampFunc = func() time.Time {
				return time.Now().UTC()
			}

			// Don't print usage on command errors.
			// PersistentPreRunE runs after argument parsing, so errors during parsing will result in printing the help
			cmd.SilenceUsage = true
			var err error

			if logFile, err = initLogging(noLogFile, logLevel, logFormat, logConsole, logFileName, logFileOverwrite); err != nil {
				return err
			}

			if os.Getenv("CQ_NO_TELEMETRY") != "" {
				return errors.New("the CQ_NO_TELEMETRY environment variable is no longer supported, unset it and set CQ_TELEMETRY_LEVEL=none instead")
			}

			oc, err := otelConfig()
			if err != nil {
				return err
			}
			if oc.Endpoint != "" {
				loggingShutdownFn, err = otel.SetupOtel(context.Background(), log.Logger, oc)
				if err != nil {
					return fmt.Errorf("failed to setup OpenTelemetry: %w", err)
				}
				log.Logger = log.Logger.Hook(otel.NewOTELLoggerHook(invocationUUID.String()))
				log.Info().Str("otel_logs_endpoint", oc.Endpoint).Msg("otel logs endpoint set")
			}

			// log version of the CLI that is running
			log.Info().Str("version", Version).Msg("CloudQuery CLI")

			sendStats := funk.ContainsString([]string{"all", "stats"}, telemetryLevel.String())
			if Version != "development" && sendStats {
				analytics.InitClient()
			}

			sendErrors := funk.ContainsString([]string{"all", "errors"}, telemetryLevel.String())
			if sentryDsn != "" && Version != "development" && sendErrors {
				if err := initSentry(sentryDsn, Version); err != nil {
					// we don't fail on sentry init errors as there might be no connection or sentry can be blocked.
					log.Warn().Err(err).Msg("failed to initialize sentry")
				}
			} else {
				disableSentry = true
			}

			return nil
		},
	}

	cmd.SetErr(secrets.NewSecretAwareWriter(os.Stderr, secretAwareRedactor))
	cmd.SetOut(secrets.NewSecretAwareWriter(os.Stdout, secretAwareRedactor))

	cmd.PersistentFlags().String("cq-dir", ".cq", "directory to store cloudquery files, such as downloaded plugins")

	// Logging Flags
	cmd.PersistentFlags().BoolVar(&logConsole, "log-console", false, "enable console logging")
	cmd.PersistentFlags().Var(logFormat, "log-format", "Logging format (json, text)")
	cmd.PersistentFlags().Var(logLevel, "log-level", "Logging level (trace, debug, info, warn, error)")
	cmd.PersistentFlags().BoolVar(&noLogFile, "no-log-file", false, "Disable logging to file")
	cmd.PersistentFlags().StringVar(&logFileName, "log-file-name", "cloudquery.log", "Log filename")
	cmd.PersistentFlags().BoolVar(&logFileOverwrite, "log-file-overwrite", false, "Overwrite log file on each run instead of appending. Use this if your filesystem does not support append mode (e.g. FUSE-mounted cloud storage).")

	// Telemetry (analytics) flags
	f := cmd.PersistentFlags().VarPF(telemetryLevel, "telemetry-level", "", "Telemetry level (none, errors, stats, all)")
	f.DefValue = "all"

	iUUID := cmd.PersistentFlags().VarPF(&invocationUUID, "invocation-id", "", "useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services")
	iUUID.DefValue = "<NEW-RANDOM-UUID>"
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})

	pluginCmd := &cobra.Command{
		Use:        "plugin",
		Short:      "Plugin commands",
		SuggestFor: []string{"install"},
	}
	pluginDocCmd := &cobra.Command{
		Use:    "docs",
		Short:  "Plugin docs commands",
		Hidden: true,
	}
	pluginDocCmd.AddCommand(
		newCmdPluginDocsDownload(),
		newCmdPluginDocsUpload(),
	)
	pluginUIAssetsCmd := &cobra.Command{
		Use:    "uiassets",
		Short:  "Plugin UI asset commands",
		Hidden: true,
	}
	pluginUIAssetsCmd.AddCommand(
		newCmdPluginUIAssetsUpload(),
	)

	pluginCmd.AddCommand(
		newCmdPluginInstall(),
		newCmdPluginPublish(),
		pluginDocCmd,
		pluginUIAssetsCmd,
	)

	addonCmd := &cobra.Command{
		Use:   "addon",
		Short: "Addon commands",
	}
	addonCmd.AddCommand(
		newCmdAddonDownload(),
		newCmdAddonPublish(),
	)

	cmd.AddCommand(
		NewCmdSync(),
		NewCmdMigrate(),
		newCmdDoc(),
		NewCmdTables(),
		newCmdLogin(),
		newCmdLogout(),
		newCmdSwitch(),
		newCmdTestConnection(),
		newCmdValidateConfig(),
		pluginCmd,
		addonCmd,
		newCmdInit(),
		newCmdUpgrade(),
	)

	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true
	cobra.OnFinalize(func() {
		analytics.Close()
		if loggingShutdownFn != nil {
			loggingShutdownFn()
		}
	})

	return cmd
}

// formats a timestamp in UTC and RFC3339
func formatTimestampUtcRfc3339(timestamp any) string {
	timestampConcrete, ok := timestamp.(time.Time)
	if !ok {
		return fmt.Sprintf("%v", timestamp)
	}

	return timestampConcrete.UTC().Format(time.RFC3339)
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}

func otelConfig() (*otel.Config, error) {
	insecureEndpoint, err := strconv.ParseBool(env.GetEnvOrDefault("OTEL_ENDPOINT_INSECURE", "false"))
	if err != nil {
		return nil, err
	}
	return &otel.Config{
		Endpoint: env.GetEnvOrDefault("OTEL_ENDPOINT", ""),
		Insecure: insecureEndpoint,
	}, nil
}
