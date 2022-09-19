package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://be7c45692567444299f8bef3de545b86@o1396617.ingest.sentry.io/6747596"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
