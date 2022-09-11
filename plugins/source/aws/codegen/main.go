package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
)

func main() {
	resources := recipes.Ec2Resources()
	resources = append(resources, recipes.AccessAnalyzerResources()...)
	resources = append(resources, recipes.EcsResources()...)
	resources = append(resources, recipes.EfsResources()...)
	resources = append(resources, recipes.EksResources()...)
	resources = append(resources, recipes.ElastiCacheResources()...)
	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}