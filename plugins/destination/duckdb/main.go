package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/client"
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = ""
)

func main() {
	p := destination.NewPlugin("duckdb", plugin.Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
