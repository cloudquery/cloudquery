package main

import (
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:                "azure",
		Provider:            resources.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
