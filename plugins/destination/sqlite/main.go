package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/client"
	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://8e769078076443cd9c578833740beb54@o1396617.ingest.sentry.io/4504797525901312"
)

func main() {
	p := destination.NewPlugin("sqlite", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
