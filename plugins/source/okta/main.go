package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://e43b6becdda446e6aedb4539cbc7cc83@o1396617.ingest.sentry.io/6747629"

func main() {
	serve.Source(plugin.Plugin(),
		serve.WithSourceSentryDSN(sentryDSN),
	)
}
