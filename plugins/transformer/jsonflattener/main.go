package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
	)
	if err := serve.Plugin(p).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
