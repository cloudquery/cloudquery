package cmd

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/ui"
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
	if client.Version == client.DevelopmentVersion && !viper.GetBool("debug-sentry") {
		dsn = "" // Disable Sentry in development mode, unless debug-sentry was enabled
	}

	if err := sentry.Init(sentry.ClientOptions{
		Debug:     viper.GetBool("debug-sentry"),
		Dsn:       dsn,
		Transport: sentrySyncTransport,
		Environment: func() string {
			if client.Version == client.DevelopmentVersion {
				return "development"
			}
			return "release"
		}(),
		Release:          "cloudquery@" + client.Version,
		AttachStacktrace: true, // send stack trace with panic recovery
		Integrations: func(it []sentry.Integration) []sentry.Integration {
			ret := make([]sentry.Integration, 0, len(it))
			for i := range it {
				switch it[i].Name() {
				case "ContextifyFrames", "Modules":
					// nothing
				default:
					ret = append(ret, it[i])
				}
			}
			return ret
		},
		ServerName: func() string {
			hn, err := os.Hostname()
			if err != nil || hn == "" {
				return "unknown" // Not returning empty string, otherwise Sentry auto-fill it
			}
			return telemetry.HashAttribute(hn)
		}(),
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint != nil && hint.RecoveredException != nil {
				// Keep stack trace on recover() events
				return event
			}

			// Remove stack trace otherwise
			for i := range event.Exception {
				event.Exception[i].Stacktrace = nil
			}

			if len(event.Exception) > 0 {
				if event.Tags["provider"] != "" {
					event.Exception[0].Type = "Diag:" + event.Tags["provider"] + "@" + event.Tags["provider_version"]
				}
			}

			return event
		},
	}); err != nil {
		zerolog.Info().Err(err).Msg("sentry.Init failed")
	}
}

func setSentryVars(traceID, randomID string) {
	if strings.HasPrefix(randomID, telemetry.CQTeamID) && !viper.GetBool("debug-sentry") {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: "",
		}); err != nil {
			zerolog.Info().Err(err).Msg("sentry.Init to disable failed")
		}
	}

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra("trace_id", traceID)
		scope.SetUser(sentry.User{
			ID: randomID,
		})
		scope.SetTags(map[string]string{
			"terminal": strconv.FormatBool(ui.IsTerminal()),
			"ci":       strconv.FormatBool(telemetry.IsCI()),
			"faas":     strconv.FormatBool(telemetry.IsFaaS()),
		})
	})
}
