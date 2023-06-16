package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://e944e0a50f1c43eaa6ffd89e7b39f4da@o1396617.ingest.sentry.io/4504734217142272"

func main() {
	serve.Plugin(client.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
}
