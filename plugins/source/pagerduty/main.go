package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/plugin"

	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://88b3936014084326bed8d93aaf24e559@o1396617.ingest.sentry.io/4504321041432576"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	err := p.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
