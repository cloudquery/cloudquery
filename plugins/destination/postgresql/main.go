//go:build !fipsEnabled

package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithConnectionTester(client.ConnectionTester),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
