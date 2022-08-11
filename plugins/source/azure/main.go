package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/provider"
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
