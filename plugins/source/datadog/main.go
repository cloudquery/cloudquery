package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://99f66f1a627f48deb66e49a25d6028a6@o1396617.ingest.sentry.io/6747628"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
