package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/gremlin/client"
	"github.com/cloudquery/cloudquery/plugins/destination/gremlin/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://eabea8e3c07b4c44a298546306641da0@o1396617.ingest.sentry.io/4504809326641152"
)

func main() {
	p := destination.NewPlugin("gremlin", plugin.Version, client.New, destination.WithManagedWriter(), destination.WithDefaultBatchSize(200))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
