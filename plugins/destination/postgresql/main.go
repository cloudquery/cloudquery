package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/resources/plugin"
	pluginSDK "github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := pluginSDK.NewPlugin(plugin.Name, plugin.Version, client.New,
		pluginSDK.WithKind(plugin.Kind),
		pluginSDK.WithTeam(plugin.Team),
		pluginSDK.WithJSONSchema(spec.JSONSchema),
		pluginSDK.WithConnectionTester(client.ConnectionTester),
	)

	if err := serve.Plugin(p, serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
