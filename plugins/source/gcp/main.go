package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/cloudquery/plugins/source/gcp/resources/plugin"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
