package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/test/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(plugin.New())
}
