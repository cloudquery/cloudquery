package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/mysql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = ""

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
