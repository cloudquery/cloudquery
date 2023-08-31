package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://c2101fe3c5fd4f91a095b2b37dc6364a@o1396617.ingest.sentry.io/4504333793165313"

func main() {
	p := serve.Plugin(plugin.Snyk(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
