package main

import (
	"context"
	"log"

	"github.com/cloudquery/plugins/vault/resources/plugin"

	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://64ac2cdbdb9cb64ed7d78053412495b8@o1396617.ingest.sentry.io/4505635677863936"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
