package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/slack/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/slack/codegen/tables"
)

func generateResources() ([]*recipes.Resource, error) {
	var resources []*recipes.Resource
	resources = append(resources, recipes.UserResources()...)
	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			return nil, err
		}
	}

	return resources, nil
}

func main() {
	resources, err := generateResources()
	if err != nil {
		log.Fatal(err)
	}

	err = tables.Generate(resources)
	if err != nil {
		log.Fatal(err)
	}
}
