package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://29efd1efe53847c4a979bc2d2e712a74@o1396617.ingest.sentry.io/4504797528981504"

func main() {
	if err := serve.Plugin(plugin.Plugin(),
		serve.WithPluginSentryDSN(sentryDSN),
	).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
