package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://e97f52f0673f44849bab5617a4e07959@o1396617.ingest.sentry.io/4504474234650624"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
