package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://9a27c2df309b49fa9520937069b370c6@o1396617.ingest.sentry.io/4504409131515904"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
