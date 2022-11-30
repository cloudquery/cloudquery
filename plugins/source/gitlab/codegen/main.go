package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/codegen/recipes"
)

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.Groups()...)
	resources = append(resources, recipes.Projects()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
	// if err := recipes.GeneratePlugin(resources); err != nil {
	// 	log.Fatal(err)
	// }
}
