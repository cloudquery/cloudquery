package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/github/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://99f66f1a627f48deb66e49a25d6028a6@o1396617.ingest.sentry.io/6747628"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
