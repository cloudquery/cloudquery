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
	p := plugin.NewPlugin("jira", internalPlugin.Version, client.New)
	if err := serve.Plugin(p).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
