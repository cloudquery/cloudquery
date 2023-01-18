package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/kinesisfirehose/client"
	"github.com/cloudquery/cloudquery/plugins/destination/kinesisfirehose/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	// TODO: Add Sentry DSN
	sentryDSN = ""
)

func main() {
	p := destination.NewPlugin("kinesisfirehose", plugin.Version, client.New, destination.WithDefaultBatchSize(500))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
