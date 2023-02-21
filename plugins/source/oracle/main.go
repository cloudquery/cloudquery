package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://e97f52f0673f44849bab5617a4e07959@o1396617.ingest.sentry.io/4504474234650624"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
