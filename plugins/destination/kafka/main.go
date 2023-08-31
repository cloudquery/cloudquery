package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/kafka/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/kafka/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

func main() {
	p := plugin.NewPlugin("kafka", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
