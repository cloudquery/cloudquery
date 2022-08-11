package main

import (
	"github.com/cloudquery/cq-provider-github/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "github",
		Provider: provider.Provider(),
	})
}
