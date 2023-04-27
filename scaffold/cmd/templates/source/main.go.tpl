package main

import (
	"github.com/cloudquery/plugin-sdk/v2/serve"
	"github.com/{{.Org}}/cq-source-{{.Name}}/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
