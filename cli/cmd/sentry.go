package cmd

import "github.com/getsentry/sentry-go"

func initSentry(sentryDsn string, version string) error {
	return sentry.Init(sentry.ClientOptions{
		Debug:      false,
		Dsn:        sentryDsn,
		Release:    "cloudquery@" + version,
		Transport:  sentry.NewHTTPSyncTransport(),
		ServerName: "oss", // set to "oss" on purpose to avoid sending any identifying information
		// https://docs.sentry.io/platforms/go/configuration/options/#removing-default-integrations
		Integrations: func(integrations []sentry.Integration) []sentry.Integration {
			var filteredIntegrations []sentry.Integration
			for _, integration := range integrations {
				if integration.Name() == "Modules" {
					continue
				}
				filteredIntegrations = append(filteredIntegrations, integration)
			}
			return filteredIntegrations
		},
	})
}
