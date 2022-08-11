package main

import (
	"github.com/cloudquery/cq-provider-azure/resources/provider"
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
