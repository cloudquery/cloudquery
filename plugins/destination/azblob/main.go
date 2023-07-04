package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/azblob/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/azblob/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://43b48b4844804de7aebffe352b044f2c@o1396617.ingest.sentry.io/4504411507392512"
)

func main() {
	p := plugin.NewPlugin("azblob", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
