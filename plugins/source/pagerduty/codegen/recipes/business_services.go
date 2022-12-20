package recipes

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BusinessServicesResources() []*Resource {
	return []*Resource{
		{
			SubService:  "business_services",
			Description: "https://developer.pagerduty.com/api-reference/e67570b9d0e3d-list-business-services",
			Struct:      pagerduty.BusinessService{},
			PKColumns:   []string{"id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "dependencies",
					Type:     schema.TypeJSON,
					Resolver: `DependenciesResolver`,
				},
			},
			Template: "paginated",
		},
	}
}
