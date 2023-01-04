package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://41d84fc1a9a64444b486c00709c099b5@o1396617.ingest.sentry.io/4504446079860736"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
