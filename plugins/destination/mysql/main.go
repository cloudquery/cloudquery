package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/mysql/v5/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/mysql/v5/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://38177cc7daaa472aa8c72496e774eda3@o1396617.ingest.us.sentry.io/4504729173950464"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
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
