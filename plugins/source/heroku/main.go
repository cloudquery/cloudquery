package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://007186e3289a490c9af043fe0f0b3fb2@o1396617.ingest.sentry.io/6765331"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
