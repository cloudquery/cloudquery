package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/codegen/resources"
)

func main() {
	err := resources.Generate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
