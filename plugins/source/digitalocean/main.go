package main

import (
	"github.com/cloudquery/cq-provider-digitalocean/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:                "digitalocean",
		Provider:            provider.Provider(),
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
