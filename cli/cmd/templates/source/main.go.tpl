package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
)

func main() {
	serve.Source(client.New())
}
