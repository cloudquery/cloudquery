package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/oracledb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://3613e006170b45399a66a93c992225e4@o1396617.ingest.sentry.io/4504922534445056"

func main() {
	pluginServe := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := pluginServe.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
