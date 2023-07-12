package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://c1a3fc67153e4f7781125a8dbc1a737f@o1396617.ingest.sentry.io/4504530915557376"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	err := p.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
