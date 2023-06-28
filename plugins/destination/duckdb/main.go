package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/duckdb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://d4bc6f4eb3014b8994c7a66846d86e18@o1396617.ingest.sentry.io/4504797281779712"
)

func main() {
	p := plugin.NewPlugin("duckdb", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN)); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
