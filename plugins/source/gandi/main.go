package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

// const sentryDSN = "https://6beffd303fbd4661a7b7c6c2d546b580@o1396617.ingest.sentry.io/6747634" TODO

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(""))
}
