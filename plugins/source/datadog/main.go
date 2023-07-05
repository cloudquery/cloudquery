package main

import (
	"context"
	"log"

	internalPlugin "github.com/cloudquery/cloudquery/plugins/source/datadog/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://a0c2570b96264748a6759bb62e8cdef5@o1396617.ingest.sentry.io/4504220208267264"

func main() {
	if err := serve.Plugin(internalPlugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
