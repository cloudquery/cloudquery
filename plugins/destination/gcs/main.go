package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/gcs/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/gcs/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://c808d26943414033b2fb8bb5b5822ab9@o1396617.ingest.sentry.io/4504407917592576"
)

func main() {
	p := plugin.NewPlugin("file", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
