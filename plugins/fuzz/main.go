package main

import (
	"github.com/cloudquery/cq-provider-fuzzing/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "fuzz",
		Provider: provider.FuzzProvider(),
	})
}
