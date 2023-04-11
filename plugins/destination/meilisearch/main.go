package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/client"
	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const (
	sentryDSN = "https://3218256ba1fb4d2fa08eb03feea5e871@o1396617.ingest.sentry.io/4504893018996736"
)

func main() {
	serve.Destination(
		destination.NewPlugin(
			"meilisearch",
			plugin.Version,
			client.New,
			destination.WithDefaultBatchSize(1000),
			destination.WithManagedWriter(),
		),
		serve.WithDestinationSentryDSN(sentryDSN),
	)
}
