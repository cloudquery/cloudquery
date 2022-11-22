package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
