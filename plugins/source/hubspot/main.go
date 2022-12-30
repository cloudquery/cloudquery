package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

//const sentryDSN = "TODO"

func main() {
	serve.Source(
		plugin.HubSpot(),
		// TODO: serve.WithSourceSentryDSN(sentryDSN),
	)
}
