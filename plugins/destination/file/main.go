package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/file/client"
	"github.com/cloudquery/cloudquery/plugins/destination/file/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://79d5e237dafe45e1a4ec0785bc528280@o1396617.ingest.sentry.io/4504083471335424"
)

func main() {
	p := plugins.NewDestinationPlugin("file", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
