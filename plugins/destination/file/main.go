package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/file/client"
	"github.com/cloudquery/cloudquery/plugins/destination/file/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://1e33dfd084aa43f2aa8e686f15a64e45@o1396617.ingest.sentry.io/4504407264526336"
)

func main() {
	p := destination.NewPlugin("file", plugin.Version, client.New, destination.WithManagedWriter())
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
