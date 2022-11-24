package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://45b253f7a0794df2bba513eaa930121b@o1396617.ingest.sentry.io/4504214828875776"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
