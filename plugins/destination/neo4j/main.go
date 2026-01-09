package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/neo4j/v5/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/neo4j/v5/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://129e24d20c7447abb3fa26d058cff048@o1396617.ingest.sentry.io/4504424944238592"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithConnectionTester(client.NewConnectionTester(client.New)),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
