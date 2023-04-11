package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const (
	sentryDSN = "https://f745aa5eaa44407ca4003a9c77a1b240@o1396617.ingest.sentry.io/4504481164754944"
)

func main() {
	serve.Destination(
		destination.NewPlugin(
			"mssql",
			plugin.Version,
			client.New,
			destination.WithDefaultBatchSize(1000),
			destination.WithManagedWriter(),
		),
		serve.WithDestinationSentryDSN(sentryDSN),
	)
}
