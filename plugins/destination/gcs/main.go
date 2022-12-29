package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/gcs/client"
	"github.com/cloudquery/cloudquery/plugins/destination/gcs/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://c808d26943414033b2fb8bb5b5822ab9@o1396617.ingest.sentry.io/4504407917592576"
)

func main() {
	p := destination.NewPlugin("gcs", plugin.Version, client.New, destination.WithManagedWriter())
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
