package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/meilisearch/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://3218256ba1fb4d2fa08eb03feea5e871@o1396617.ingest.sentry.io/4504893018996736"
)

func main() {
	serve.Plugin(
		plugin.NewPlugin(
			"meilisearch",
			internalPlugin.Version,
			client.New,
		),
		serve.WithPluginSentryDSN(sentryDSN),
	)
}
