package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://e944e0a50f1c43eaa6ffd89e7b39f4da@o1396617.ingest.sentry.io/4504734217142272"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	err := p.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
