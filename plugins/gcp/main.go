package main

import (
	"github.com/cloudquery/cloudquery/plugins/gcp/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "gcp",
		Provider: provider.Provider(),
	})
}
