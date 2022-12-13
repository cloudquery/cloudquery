package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/codegen/recipes"
)

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.Core()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
	if err := recipes.GenerateTablesList(resources); err != nil {
		log.Fatal(err)
	}
}
