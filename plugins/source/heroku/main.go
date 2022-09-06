package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "heroku",
		Provider: provider.Provider(),
	})
}
