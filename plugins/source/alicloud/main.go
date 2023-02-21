package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://c1a3fc67153e4f7781125a8dbc1a737f@o1396617.ingest.sentry.io/4504530915557376"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
