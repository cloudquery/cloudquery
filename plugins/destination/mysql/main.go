package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/mysql/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/mysql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://38177cc7daaa472aa8c72496e774eda3@o1396617.ingest.sentry.io/4504729173950464"
)

func main() {
	if err := serve.Plugin(
		plugin.NewPlugin(
			"mysql",
			internalPlugin.Version,
			client.New,
		),
		serve.WithDestinationV0V1Server(),
		serve.WithPluginSentryDSN(sentryDSN),
	).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
