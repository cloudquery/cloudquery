package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/alicloud/codegen/recipies"
)

func main() {
	for _, f := range []func() []*recipies.Resource{
		recipies.BSS,
		recipies.OSS,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
