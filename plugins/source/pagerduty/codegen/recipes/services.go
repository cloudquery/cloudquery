package recipes

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServiceResources() []*Resource {
	return []*Resource{
		{
			SubService:                    "services",
			PKColumns:                     []string{"id"},
			Struct:                        pagerduty.Service{},
			ListOptionsStructNameOverride: "ListServiceOptions",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "dependencies",
					Type:     schema.TypeJSON,
					Resolver: `DependenciesResolver`,
				},
			},
			Description: "https://developer.pagerduty.com/api-reference/e960cca205c0f-list-services",

			Relations: []string{
				"ServiceRules()",
			},

			ResponseStructOverride: "ListServiceResponse",
			SkipMockGeneration:     true,
		},
		{
			SubService:  "service_rules",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.ServiceRule{},
			Description: "https://developer.pagerduty.com/api-reference/d69ad7f1ec900-list-service-s-event-rules",

			Template: "nested",

			SkipFetchGeneration: true,

			ResponseFieldOverride: "Rules",
			RestPathOverride:      "/rules",
		},
	}
}
