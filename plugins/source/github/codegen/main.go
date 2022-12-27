package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/github/codegen/recipes"
)

func main() {
	for _, f := range []func() []*recipes.Resource{
		recipes.Actions,
		recipes.Billing,
		recipes.External,
		recipes.Hooks,
		recipes.Installations,
		recipes.Issues,
		recipes.Organizations,
		recipes.Repositories,
		recipes.Teams,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
