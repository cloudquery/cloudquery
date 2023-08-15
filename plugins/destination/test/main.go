package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/test/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin("test", internalPlugin.Version, client.New)
	if err := serve.Plugin(p,
		serve.WithDestinationV0V1Server(),
	).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
