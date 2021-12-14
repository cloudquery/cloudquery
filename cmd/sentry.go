package cmd

import (
	"time"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/getsentry/sentry-go"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultSentryFlushTimeout = 2 * time.Second

func registerSentryFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("sentry-debug", false, "Enable Sentry APIDebug")
	cmd.PersistentFlags().String("sentry-dsn", "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")
	// cmd.PersistentFlags().String("sentry-environment", "development", "Sentry Environment")
	cmd.PersistentFlags().String("sentry-release", client.Version, "Sentry Release")

	_ = viper.BindPFlag("sentry-debug", cmd.PersistentFlags().Lookup("sentry-debug"))
	_ = viper.BindPFlag("sentry-dsn", cmd.PersistentFlags().Lookup("sentry-dsn"))
	// _ = viper.BindPFlag("sentry-environment", cmd.PersistentFlags().Lookup("sentry-environment"))
	_ = viper.BindPFlag("sentry-release", cmd.PersistentFlags().Lookup("sentry-release"))
}

func initSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Debug: viper.GetBool("sentry-debug"),
		Dsn:   viper.GetString("sentry-dsn"),
		// Environment: viper.GetString("sentry-environment"),
		Release: viper.GetString("sentry-release"),
	}); err != nil {
		zerolog.Info().Err(err).Msg("sentry.Init failed")
	}
}

func flushSentry(_ *cobra.Command, _ []string) {
	sentry.Flush(defaultSentryFlushTimeout)
}
