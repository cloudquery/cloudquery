package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
