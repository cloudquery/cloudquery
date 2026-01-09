package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/client"
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://03ef940ae6064ae8a69fa5e1ad74d3f2@o1396617.ingest.sentry.io/4504600056299520"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithConnectionTester(client.NewConnectionTester(client.New)),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
