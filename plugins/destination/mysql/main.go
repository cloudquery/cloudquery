package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/mysql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/mysql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://38177cc7daaa472aa8c72496e774eda3@o1396617.ingest.sentry.io/4504729173950464"
)

func main() {
	serve.Destination(
		destination.NewPlugin(
			"mysql",
			plugin.Version,
			client.New,
			destination.WithManagedWriter(),
		),
		serve.WithDestinationSentryDSN(sentryDSN),
	)
}
