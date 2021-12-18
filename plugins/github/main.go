package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	// CHANGEME: change this to your package name
	"github.com/cloudquery/cq-provider-template/resources"
)

func main() {
	p := resources.Provider()
	serve.Serve(&serve.Options{
		Name:                p.Name,
		Provider:            p,
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
