package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/jira/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin("jira", "0.0.1", client.New)
	if err := serve.Plugin(p).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
