package main

import (
	"context"
	"log"

	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/cloudquery/plugins/source/gcp/resources/plugin"
)

const sentryDSN = "https://c30e57a331fe4101a11b3c83d780793f@o1396617.ingest.sentry.io/6720365"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
