package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
)

func main() {
	resources := recipes.Ec2Resources()
	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}