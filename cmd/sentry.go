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
	cmd.PersistentFlags().Bool("no-sentry", false, "Disable Sentry")
	cmd.PersistentFlags().Bool("debug-sentry", false, "Enable Sentry debug mode")
	cmd.PersistentFlags().String("sentry-dsn", "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	_ = cmd.PersistentFlags().MarkHidden("sentry-dsn")

	_ = viper.BindPFlag("no-sentry", cmd.PersistentFlags().Lookup("no-sentry"))
	_ = viper.BindPFlag("debug-sentry", cmd.PersistentFlags().Lookup("debug-sentry"))
	_ = viper.BindPFlag("sentry-dsn", cmd.PersistentFlags().Lookup("sentry-dsn"))
}

func initSentry() {
	dsn := viper.GetString("sentry-dsn")
	if viper.GetBool("no-sentry") {
		dsn = "" // "To drop all events, set the DSN to the empty string."
	}

	if err := sentry.Init(sentry.ClientOptions{
		Debug:   viper.GetBool("debug-sentry"),
		Dsn:     dsn,
		Release: client.Version,
	}); err != nil {
		zerolog.Info().Err(err).Msg("sentry.Init failed")
	}
}

func flushSentry(_ *cobra.Command, _ []string) {
	sentry.Flush(defaultSentryFlushTimeout)
}
