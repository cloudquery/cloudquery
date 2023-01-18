package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/azblob/client"
	"github.com/cloudquery/cloudquery/plugins/destination/azblob/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://43b48b4844804de7aebffe352b044f2c@o1396617.ingest.sentry.io/4504411507392512"
)

func main() {
	p := destination.NewPlugin("azblob", plugin.Version, client.New, destination.WithManagedWriter())
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
