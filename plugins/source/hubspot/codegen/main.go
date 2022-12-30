package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/hubspot/codegen/resources"
)

func main() {
	err := resources.Generate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
