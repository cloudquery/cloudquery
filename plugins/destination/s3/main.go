package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/s3/client"
	"github.com/cloudquery/cloudquery/plugins/destination/s3/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/s3/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	server := serve.Plugin(p, serve.WithDestinationV0V1Server())

	done := instrumentPprof()
	defer done()
	err := server.Serve(context.Background())
	if err != nil {
		log.Println("failed to serve plugin:", err)
	}
}
