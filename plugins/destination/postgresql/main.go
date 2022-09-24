package main

import (
	"github.com/cloudquery/cloudquery/plugins/destinations/postgresql/client"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(client.New())
}

