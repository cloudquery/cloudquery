package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDsn = "https://a0d398d362d34e63a7ed246f7cb76b5a@o1396617.ingest.sentry.io/4504798123786240"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDsn)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
