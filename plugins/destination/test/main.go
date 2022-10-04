package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/test/client"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

var (
	Version = "Development"
)

func main() {
	serve.Destination(plugins.NewDestinationPlugin("test", Version, client.New))
}
