package main

import (
	"github.com/cloudquery/cq-provider-k8s/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	provider := resources.Provider()
	serve.Serve(&serve.Options{
		Name:     provider.Name,
		Provider: provider,
	})
}
