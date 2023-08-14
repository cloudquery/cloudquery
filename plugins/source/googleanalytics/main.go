package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://1389ab4b050f4862879db19a783ecb78@o1396617.ingest.sentry.io/4504792343379968"

func main() {
	if err := serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN)).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
