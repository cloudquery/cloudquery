package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = ""
)

func main() {
	p := plugins.NewDestinationPlugin("sqlite", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
