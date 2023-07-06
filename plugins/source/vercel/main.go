package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://236c9fd8bbe247ab9a10901438d57641@o1396617.ingest.sentry.io/4504316330901504"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
