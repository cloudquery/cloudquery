package main

import (
	"github.com/cloudquery/cq-provider-fuzz/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "fuzz",
		Provider: provider.FuzzProvider(),
	})
}
