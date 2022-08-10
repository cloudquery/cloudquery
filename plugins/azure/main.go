package main

import (
	"github.com/cloudquery/cloudquery/plugins/azure/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:                "azure",
		Provider:            provider.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
