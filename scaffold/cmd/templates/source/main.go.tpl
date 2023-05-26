package main

import (
	"github.com/{{.Org}}/cq-source-{{.Name}}/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
