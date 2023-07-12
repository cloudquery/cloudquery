package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/awspricing/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://a7dd3cdeae17424d8cd70f9c76c38fe1@o1396617.ingest.sentry.io/4504724922368000"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
