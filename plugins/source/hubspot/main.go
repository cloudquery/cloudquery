package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDsn = "https://015fd88288884743b76b50d9dfc14130@o1396617.ingest.sentry.io/4504559739011072"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDsn)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
