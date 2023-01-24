package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/plausible/resources/plugin"

	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDsn = "https://1a277b8d8bcb49148c5838bcc01de188@o1396617.ingest.sentry.io/4504548722606080"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDsn))
}
