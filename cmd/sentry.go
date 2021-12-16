package cmd

import (
	"time"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/getsentry/sentry-go"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerSentryFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("debug-sentry", false, "Enable Sentry debug mode")
	cmd.PersistentFlags().String("sentry-dsn", "https://5ff9e378a79d4ba2821f540b036286e9@o912044.ingest.sentry.io/6106324", "Sentry DSN")

	_ = cmd.PersistentFlags().MarkHidden("sentry-dsn")

	_ = viper.BindPFlag("debug-sentry", cmd.PersistentFlags().Lookup("debug-sentry"))
	_ = viper.BindPFlag("sentry-dsn", cmd.PersistentFlags().Lookup("sentry-dsn"))
}

func initSentry() {
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 2

	dsn := viper.GetString("sentry-dsn")
	if viper.GetBool("no-telemetry") {
		dsn = "" // "To drop all events, set the DSN to the empty string."
	}

	if err := sentry.Init(sentry.ClientOptions{
		Debug:     viper.GetBool("debug-sentry"),
		Dsn:       dsn,
		Transport: sentrySyncTransport,
		Environment: func() string {
			if client.Version == client.DefaultVersion {
				return "development"
			}
			return "release"
		}(),
		Release:          client.Version,
		AttachStacktrace: true,
	}); err != nil {
		zerolog.Info().Err(err).Msg("sentry.Init failed")
	}
}
