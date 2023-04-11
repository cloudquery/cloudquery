package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/test/client"
	"github.com/cloudquery/cloudquery/plugins/destination/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/serve"
)

func main() {
	serve.Destination(destination.NewPlugin("test", plugin.Version, client.New))
}
