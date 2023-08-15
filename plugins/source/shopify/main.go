package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://9a27c2df309b49fa9520937069b370c6@o1396617.ingest.sentry.io/4504409131515904"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
