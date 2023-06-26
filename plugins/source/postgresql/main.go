package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/postgresql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://995c68a7e67541338e22dd8120e81c42@o1396617.ingest.sentry.io/4504316028452864"

func main() {
	pluginServe := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := pluginServe.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
