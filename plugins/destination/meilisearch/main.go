package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://3218256ba1fb4d2fa08eb03feea5e871@o1396617.ingest.sentry.io/4504893018996736"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
