package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://ae63e383a6724d1cad4bf9546c7d5489@o1396617.ingest.sentry.io/4504560038838272"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
