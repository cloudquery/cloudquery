package main

import (
	"github.com/cloudquery/cloudquery/plugins/aws/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "aws",
		Provider: provider.Provider(),
	})
}
