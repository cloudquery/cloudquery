package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/codegen/resources"
)

func main() {
	err := resources.Generate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
