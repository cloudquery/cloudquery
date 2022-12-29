package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

func main() {
	p := destination.NewPlugin("postgresql", plugin.Version, client.New, destination.WithDefaultBatchSize(1000))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
