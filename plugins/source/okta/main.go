package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
