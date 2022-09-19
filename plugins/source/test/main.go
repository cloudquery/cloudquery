package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
