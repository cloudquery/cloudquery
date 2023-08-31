package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://4606a1601b52488889313375ba91df89@o1396617.ingest.sentry.io/4504441433751552"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
