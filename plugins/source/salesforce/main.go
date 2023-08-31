package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://77af2b07ee5c45519f17f2f6314421c4@o1396617.ingest.sentry.io/4504390800310272"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
