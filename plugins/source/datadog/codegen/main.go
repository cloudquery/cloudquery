package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/codegen/services"
)

func main() {
	err := services.Generate()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range []func() []*recipes.Resource{
		recipes.DashboardLists,
		recipes.Dashboards,
		recipes.Downtimes,
		recipes.Hosts,
		recipes.Incidents,
		recipes.Monitors,
		recipes.Notebooks,
		recipes.Roles,
		recipes.Synthetics,
		recipes.Users,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
