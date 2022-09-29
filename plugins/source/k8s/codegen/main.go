package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/codegen/recipes"
)

func main() {
	resources := recipes.CoreResources()
	resources = append(resources, recipes.AppsResources()...)
	resources = append(resources, recipes.BatchResources()...)
	resources = append(resources, recipes.NetworkingResources()...)
	resources = append(resources, recipes.RbacResources()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
