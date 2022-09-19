package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
