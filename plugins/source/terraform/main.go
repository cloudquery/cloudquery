package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "terraform",
		Provider: provider.Provider(),
	})
}
