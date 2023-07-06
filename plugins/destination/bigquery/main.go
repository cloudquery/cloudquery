package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/bigquery/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/bigquery/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://8856f7c90f284b0f912f5873a6448ca3@o1396617.ingest.sentry.io/4504220665577472"
)

func main() {
	p := plugin.NewPlugin("bigquery", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithDestinationV0V1Server(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
