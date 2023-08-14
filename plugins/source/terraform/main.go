package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://8a66aa7c550b46258f2391cbae261fe2@o1396617.ingest.sentry.io/6747630"

func main() {
	p := serve.Plugin(plugin.Terraform(), serve.WithPluginSentryDSN(sentryDSN))
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
