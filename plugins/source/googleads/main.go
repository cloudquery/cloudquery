package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const sentryDSN = "https://5f6d36d7025942e28c48dac425a0d09f@o1396617.ingest.sentry.io/4504881508646912"

func main() {
	serve.Source(plugin.Plugin(),
		serve.WithSourceSentryDSN(sentryDSN),
	)
}
