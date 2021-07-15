package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/cloudquery/cq-provider-template/resources"
)

func main() {
	serve.Serve(&serve.Options{
		Name:                "aws",
		Provider:            resources.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
