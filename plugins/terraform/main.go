package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/cloudquery/cq-provider-terraform/resources"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "terraform",
		Provider: resources.Provider(),
	})
}
