package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/salesforce/resources/plugin"

	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDsn = "https://77af2b07ee5c45519f17f2f6314421c4@o1396617.ingest.sentry.io/4504390800310272"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDsn))
}
