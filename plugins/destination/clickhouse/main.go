package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
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
	if err := serve.Plugin(p, serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
