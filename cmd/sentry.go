package cmd

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/helpers/limit"
	"github.com/getsentry/sentry-go"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initSentry(rootOptions) {
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 2

	dsn := viper.GetString("sentry-dsn")
	if viper.GetBool("no-telemetry") {
		dsn = "" // "To drop all events, set the DSN to the empty string."
	}
	if core.Version == core.DevelopmentVersion && !viper.GetBool("sentry-debug") {
		dsn = "" // Disable Sentry in development mode, unless debug-sentry was enabled
	}
	userId := analytics.GetCookieId()

	if err := sentry.Init(sentry.ClientOptions{
		Debug:     viper.GetBool("sentry-debug"),
		Dsn:       dsn,
		Transport: sentrySyncTransport,
		Environment: func() string {
			if core.Version == core.DevelopmentVersion {
				return "development"
			}
			return "release"
		}(),
		Release:          "cloudquery@" + core.Version,
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
			return analytics.HashAttribute(hn)
		}(),
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if event.Tags["provider"] != "" {
				// Save core version in separate tag and report provider version as Release
				event.Tags["core_version"] = event.Release
				event.Release = event.Tags["provider"] + "@" + strings.TrimPrefix(event.Tags["provider_version"], "v")
			}

			if len(event.Exception) > 0 && event.Tags["provider"] != "" {
				event.Exception[0].Type = "Diag:" + event.Tags["provider"] + "@" + event.Tags["provider_version"]
			}

			if hint != nil && hint.RecoveredException != nil {
				// Keep stack trace on recover() events
				return event
			}

			// Remove stack trace otherwise
			for i := range event.Exception {
				event.Exception[i].Stacktrace = nil
			}

			return event
		},
	}); err != nil {
		zerolog.Info().Err(err).Msg("sentry.Init failed")
	}
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{
			ID: userId.String(),
		})
		scope.SetTags(map[string]string{
			"terminal":    strconv.FormatBool(ui.IsTerminal()),
			"ci":          strconv.FormatBool(analytics.IsCI()),
			"faas":        strconv.FormatBool(analytics.IsFaaS()),
			"instance_id": instanceId.String(),
		})
		scope.SetExtra("cookie_id", userId.String())
		scope.SetExtra("goroutine_count", runtime.NumGoroutine())
		ulimit, err := limit.GetUlimit()
		if err == nil && ulimit.Max != 0 {
			scope.SetExtra("current_ulimit", ulimit.Cur)
			scope.SetExtra("max_ulimit", ulimit.Max)
		}
	})
}
