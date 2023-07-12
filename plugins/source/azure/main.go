package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://be7c45692567444299f8bef3de545b86@o1396617.ingest.sentry.io/6747596"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
