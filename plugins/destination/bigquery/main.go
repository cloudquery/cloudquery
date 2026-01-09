//go:build !fipsEnabled

package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/bigquery/v4/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/bigquery/v4/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://8856f7c90f284b0f912f5873a6448ca3@o1396617.ingest.sentry.io/4504220665577472"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithBuildTargets(plugin.DefaultBuildTargets),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithConnectionTester(client.TestConnection),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
