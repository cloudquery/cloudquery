package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/test/client"
	"github.com/cloudquery/cloudquery/plugins/destination/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(destination.NewPlugin("test", plugin.Version, client.New))
}
