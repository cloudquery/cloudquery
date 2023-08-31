package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/mysql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://128084695e12456cb6b9be9b32f097da@o1396617.ingest.sentry.io/4504830878416896"

func main() {
	pluginServe := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := pluginServe.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
