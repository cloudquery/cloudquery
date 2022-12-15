package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/codegen/recipes"
)

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.Groups()...)
	resources = append(resources, recipes.Projects()...)
	resources = append(resources, recipes.Settings()...)
	resources = append(resources, recipes.Users()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
