package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/bigquery/client"
	"github.com/cloudquery/cloudquery/plugins/destination/bigquery/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://8856f7c90f284b0f912f5873a6448ca3@o1396617.ingest.sentry.io/4504220665577472"
)

func main() {
	p := destination.NewPlugin("bigquery", plugin.Version, client.New, destination.WithManagedWriter(), destination.WithDefaultBatchSize(1000))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
