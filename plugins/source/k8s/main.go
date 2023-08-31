package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

var sentryDsn = `https://f13561aecb1e42508070ef5743feb6fe@o1396617.ingest.sentry.io/6771378`

func main() {
	if err := serve.Plugin(
		plugin.Plugin(),
		serve.WithPluginSentryDSN(sentryDsn),
	).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
