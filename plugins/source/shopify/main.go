package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "" // TODO

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
