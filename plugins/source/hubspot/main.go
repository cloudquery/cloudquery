package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDsn = "https://015fd88288884743b76b50d9dfc14130@o1396617.ingest.sentry.io/4504559739011072"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDsn))
}
