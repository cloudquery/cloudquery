package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://e43b6becdda446e6aedb4539cbc7cc83@o1396617.ingest.sentry.io/6747629"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
