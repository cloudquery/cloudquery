package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://c84623f95f2f48de9562fd819d1fa78c@o1396617.ingest.sentry.io/6747621"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
