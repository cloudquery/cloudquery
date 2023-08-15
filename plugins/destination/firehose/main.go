package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/firehose/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/firehose/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://03ef940ae6064ae8a69fa5e1ad74d3f2@o1396617.ingest.sentry.io/4504600056299520"
)

func main() {
	p := plugin.NewPlugin("firehose", internalPlugin.Version, client.New)
	if err := serve.Plugin(p,
		serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server(),
	).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
