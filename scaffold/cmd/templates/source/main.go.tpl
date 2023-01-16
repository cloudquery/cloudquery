package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/{{.Org}}/cq-source-{{.Name}}/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
