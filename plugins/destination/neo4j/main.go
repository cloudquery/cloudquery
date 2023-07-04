package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/neo4j/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/neo4j/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://129e24d20c7447abb3fa26d058cff048@o1396617.ingest.sentry.io/4504424944238592"
)

func main() {
	p := plugin.NewPlugin("neo4j", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
