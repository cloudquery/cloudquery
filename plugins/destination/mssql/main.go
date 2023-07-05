package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://f745aa5eaa44407ca4003a9c77a1b240@o1396617.ingest.sentry.io/4504481164754944"
)

func main() {
	p := plugin.NewPlugin("mssql", internalPlugin.Version, client.New)
	if err := serve.Plugin(p,
		serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server(),
	).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
