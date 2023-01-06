package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
