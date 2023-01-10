package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://d6d0882d63ba412c8fa88f3d3722a9d9@o1396617.ingest.sentry.io/4504322445017088"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
