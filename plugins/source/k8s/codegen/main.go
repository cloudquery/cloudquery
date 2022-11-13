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
	resources = append(resources, recipes.AdmissionRegistration()...)
	resources = append(resources, recipes.Autoscaling()...)
	resources = append(resources, recipes.Certificates()...)
	resources = append(resources, recipes.Coordination()...)
	resources = append(resources, recipes.Discovery()...)
	resources = append(resources, recipes.Nodes()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
