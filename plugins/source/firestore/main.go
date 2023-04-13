package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

const sentryDSN = "" // TODO

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
