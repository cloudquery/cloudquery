package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/client"
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	// TODO: Add Sentry DSN
	sentryDSN = ""
)

func main() {
	p := destination.NewPlugin("firehose", plugin.Version, client.New, destination.WithDefaultBatchSize(500))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
