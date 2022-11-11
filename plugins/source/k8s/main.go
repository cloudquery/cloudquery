package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/plugin"

	"github.com/cloudquery/plugin-sdk/serve"
)

var sentryDsn = `https://f13561aecb1e42508070ef5743feb6fe@o1396617.ingest.sentry.io/6771378`

func main() {
	serve.Source(
		plugin.Plugin(),
		serve.WithSourceSentryDSN(sentryDsn),
	)
}
