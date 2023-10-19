package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/jira/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/source/jira/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithKind(internalPlugin.Kind),
	)
	if err := serve.Plugin(p).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
