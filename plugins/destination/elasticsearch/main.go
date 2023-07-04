package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/elasticsearch/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/elasticsearch/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://34a198627e814d16849aeac61134f8f6@o1396617.ingest.sentry.io/4504598918922240"
)

func main() {
	p := plugin.NewPlugin("elasticsearch", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
