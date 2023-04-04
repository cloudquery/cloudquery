package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://128084695e12456cb6b9be9b32f097da@o1396617.ingest.sentry.io/4504830878416896"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
