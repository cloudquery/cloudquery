package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/client"
	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = ""
)

func main() {
	p := plugins.NewDestinationPlugin("snowflake", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
