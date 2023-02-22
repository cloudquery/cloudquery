package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://a7dd3cdeae17424d8cd70f9c76c38fe1@o1396617.ingest.sentry.io/4504724922368000"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
