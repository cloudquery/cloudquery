package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/github/gen/recipies"
)

func main() {
	for _, f := range []func() []*recipies.Resource{
		recipies.Billing,
		recipies.External,
		recipies.Hooks,
		recipies.Installations,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
