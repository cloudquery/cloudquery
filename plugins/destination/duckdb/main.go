package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/client"
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const (
	sentryDSN = "https://d4bc6f4eb3014b8994c7a66846d86e18@o1396617.ingest.sentry.io/4504797281779712"
)

func main() {
	p := destination.NewPlugin("duckdb", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
