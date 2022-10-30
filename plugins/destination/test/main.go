package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/test/client"
	"github.com/cloudquery/cloudquery/plugins/destination/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(plugins.NewDestinationPlugin("test", plugin.Version, client.New))
}
