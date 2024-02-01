package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/mongodb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://f4351c889aae4dde9b87a916d44ce836@o1396617.ingest.sentry.io/4504374724657152"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
