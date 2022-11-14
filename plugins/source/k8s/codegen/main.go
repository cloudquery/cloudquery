package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/codegen/recipes"
)

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.Discovery()...)
	resources = append(resources, recipes.AdmissionRegistration()...)
	resources = append(resources, recipes.Apps()...)
	resources = append(resources, recipes.Autoscaling()...)
	resources = append(resources, recipes.Batch()...)
	resources = append(resources, recipes.Certificates()...)
	resources = append(resources, recipes.Coordination()...)
	resources = append(resources, recipes.Core()...)
	resources = append(resources, recipes.Networking()...)
	resources = append(resources, recipes.Nodes()...)
	resources = append(resources, recipes.Rbac()...)
	resources = append(resources, recipes.Storage()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
	if err := recipes.GeneratePlugin(resources); err != nil {
		log.Fatal(err)
	}
}

