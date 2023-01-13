package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/alicloud/codegen/recipes"
)

func main() {
	for _, f := range []func() []*recipes.Resource{
		recipes.BSS,
		recipes.OSS,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
