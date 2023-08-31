package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/gremlin/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/gremlin/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://eabea8e3c07b4c44a298546306641da0@o1396617.ingest.sentry.io/4504809326641152"
)

func main() {
	p := plugin.NewPlugin("gremlin", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
