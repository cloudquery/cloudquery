package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/client"
	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = ""
)

func main() {
	p := destination.NewPlugin("sqlite", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
