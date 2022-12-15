package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/codegen/recipes"
)

func main() {
	for _, resource := range recipes.Resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
	if err := recipes.GenerateTablesList(recipes.Resources); err != nil {
		log.Fatal(err)
	}
}
