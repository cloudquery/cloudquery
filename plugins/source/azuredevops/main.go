package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://81e8135dfd0e42629810a434e0dd72cb@o1396617.ingest.sentry.io/4504317272915968"

func main() {
	p := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
