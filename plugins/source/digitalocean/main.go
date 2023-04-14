package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const sentryDSN = "https://c84623f95f2f48de9562fd819d1fa78c@o1396617.ingest.sentry.io/6747621"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
