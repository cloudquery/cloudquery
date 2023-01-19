package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://4606a1601b52488889313375ba91df89@o1396617.ingest.sentry.io/4504441433751552"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
