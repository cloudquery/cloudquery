package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
