package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(client.New())
}
