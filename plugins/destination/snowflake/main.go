package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/client"
	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://6640bc34f8d34a3d816f53d396fe997e@o1396617.ingest.sentry.io/4504208023224320"
)

func main() {
	p := plugins.NewDestinationPlugin("snowflake", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
