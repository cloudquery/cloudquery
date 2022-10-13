package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/github/codegen/recipies"
)

func main() {
	for _, f := range []func() []*recipies.Resource{
		recipies.Billing,
		recipies.External,
		recipies.Hooks,
		recipies.Installations,
		recipies.Issues,
		recipies.Organizations,
		recipies.Repositories,
		recipies.Teams,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
