package main

import "github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"

func main() {
	resources := recipes.GlueResources()
	for _, resource := range resources {
		resource.Generate()
	}
}