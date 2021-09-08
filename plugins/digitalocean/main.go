package main

import (
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:                "digitalocean",
		Provider:            resources.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
