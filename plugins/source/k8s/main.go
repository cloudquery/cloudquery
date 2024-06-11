package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	if err := serve.Plugin(plugin.Plugin()).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
