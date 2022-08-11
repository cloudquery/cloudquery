package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "terraform",
		Provider: resources.Provider(),
	})
}
