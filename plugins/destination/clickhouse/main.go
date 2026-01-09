package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/client"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://f4e38a5ddc3b49ee9e0c6a7f1bce68e0@o1396617.ingest.sentry.io/4504491617878016"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithConnectionTester(client.NewConnectionTester(client.New)),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
