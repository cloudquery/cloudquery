package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://236c9fd8bbe247ab9a10901438d57641@o1396617.ingest.sentry.io/4504316330901504"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
