package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/{{.Org}}/cq-destination-{{.Name}}/client"
)

func main() {
	serve.Destination(client.New())
}
