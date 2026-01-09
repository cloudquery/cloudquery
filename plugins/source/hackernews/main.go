package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://8e769078076443cd9c578833740beb54@o1396617.ingest.sentry.io/4504797525901312"
)

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
