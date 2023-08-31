package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/resources/plugin"
	pluginSDK "github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

func main() {
	p := pluginSDK.NewPlugin("postgresql", plugin.Version, client.New)
	server := serve.Plugin(p,
		serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server(),
	)
	err := server.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
