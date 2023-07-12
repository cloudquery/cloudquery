package main

import (
	"context"
	"log"

	internalPlugin "github.com/cloudquery/cloudquery/plugins/source/aws/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://6c6b72bc946844cb8471f49eba485cde@o1396617.ingest.sentry.io/6747636"

func main() {
	if err := serve.Plugin(internalPlugin.AWS(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
