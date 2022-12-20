package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/resources/plugin"

	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDsn = "https://88b3936014084326bed8d93aaf24e559@o1396617.ingest.sentry.io/4504321041432576"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDsn))
}
