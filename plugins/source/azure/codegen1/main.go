package main

import (
	"fmt"
	"reflect"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen1/recipes"
)

type Resource struct {
	Service string
	SubService string
	ResponseStruct string
	Client string
	ListFunc string
	NewFunc string
}

func main() {
	var resources []*recipes.Resource	
	resources = append(resources, recipes.Armadvisor()...)
	for _, r := range resources {
		v := reflect.TypeOf(r.NewFunc)
		results := v.Out(0)
		fmt.Println(results.Name())
	}	
}