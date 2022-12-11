package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/codegen/recipes"
)

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.CrowdScore()...)
	resources = append(resources, recipes.Alerts()...)

	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
