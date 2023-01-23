package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/client"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://f4351c889aae4dde9b87a916d44ce836@o1396617.ingest.sentry.io/4504374724657152"
)

func main() {
	p := destination.NewPlugin("mongodb", plugin.Version, client.New, destination.WithManagedWriter())
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
