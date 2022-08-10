package main

import (
	"github.com/cloudquery/cloudquery/plugins/cloudflare/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "cloudflare",
		Provider: provider.Provider(),
	})
}
