package main

import (
	"github.com/cloudquery/cq-provider-k8s/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	provider := provider.Provider()
	serve.Serve(&serve.Options{
		Name:     provider.Name,
		Provider: provider,
	})
}
