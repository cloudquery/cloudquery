package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/csv/client"
	"github.com/cloudquery/cloudquery/plugins/destination/csv/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	// TODO: add sentry DSN
	sentryDSN = ""
)

func main() {
	p := plugins.NewDestinationPlugin("csv", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
