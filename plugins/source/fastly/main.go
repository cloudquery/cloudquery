package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://c4e711b87a8846de90e5ba2c785fb901@o1396617.ingest.sentry.io/4504379395604480"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
