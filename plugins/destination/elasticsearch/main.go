package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/elasticsearch/client"
	"github.com/cloudquery/cloudquery/plugins/destination/elasticsearch/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v3/serve"
)

const (
	sentryDSN = "https://34a198627e814d16849aeac61134f8f6@o1396617.ingest.sentry.io/4504598918922240"
)

func main() {
	p := destination.NewPlugin("elasticsearch", plugin.Version, client.New,
		destination.WithManagedWriter(),
	)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
