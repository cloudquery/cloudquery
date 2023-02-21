package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://c44156024ef14209ad89e129e44a4a8f@o1396617.ingest.sentry.io/4504524863373312"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
