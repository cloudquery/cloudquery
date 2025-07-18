package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithConnectionTester(client.NewConnectionTester(client.New)),
	)

	if err := serve.Plugin(p, serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
