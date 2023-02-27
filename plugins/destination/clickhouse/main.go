package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/client"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://f4e38a5ddc3b49ee9e0c6a7f1bce68e0@o1396617.ingest.sentry.io/4504491617878016"
)

func main() {
	serve.Destination(
		destination.NewPlugin(
			"clickhouse",
			plugin.Version,
			client.New,
			destination.WithDefaultBatchSize(10000),
			destination.WithManagedWriter(),
		),
		serve.WithDestinationSentryDSN(sentryDSN),
	)
}
