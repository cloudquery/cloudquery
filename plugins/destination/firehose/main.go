package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/client"
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://03ef940ae6064ae8a69fa5e1ad74d3f2@o1396617.ingest.sentry.io/4504600056299520"
)

func main() {
	p := destination.NewPlugin("firehose", plugin.Version, client.New, destination.WithDefaultBatchSize(500))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
