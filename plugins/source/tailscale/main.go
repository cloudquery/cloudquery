package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

// TODO: fill in proper DSN
const sentryDSN = "TODO"

func main() {
	serve.Source(plugin.Tailscale())
	// TODO: use sentryDSN
	//serve.Source(plugin.Tailscale(), serve.WithSourceSentryDSN(sentryDSN))
}
