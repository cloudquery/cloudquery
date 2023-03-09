package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/gremlin/client"
	"github.com/cloudquery/cloudquery/plugins/destination/gremlin/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "" // TODO
)

func main() {
	p := destination.NewPlugin("gremlin", plugin.Version, client.New, destination.WithManagedWriter(), destination.WithDefaultBatchSize(1000))
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
