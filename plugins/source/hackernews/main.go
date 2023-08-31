package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://41d84fc1a9a64444b486c00709c099b5@o1396617.ingest.sentry.io/4504446079860736"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	err := p.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
