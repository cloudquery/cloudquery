package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StatsResources() []*Resource {
	resources := []*Resource{
		//{
		//	TableName:   "stats_usage",
		//	DataStruct:  &fastly.RegionsUsage{},
		//	Description: "https://developer.fastly.com/reference/api/metrics-stats/historical-stats/#get-usage",
		//	PKColumns:   []string{"id"},
		//	SkipFields:  []string{},
		//	Relations:   []string{},
		//},
		//{
		//	TableName:   "stats_usage_by_service",
		//	DataStruct:  &fastly.ServicesByRegionsUsage{},
		//	Description: "https://developer.fastly.com/reference/api/metrics-stats/historical-stats/#get-usage",
		//	PKColumns:   []string{},
		//	SkipFields:  []string{},
		//	Relations:   []string{},
		//},
		//{
		//	TableName:   "stats_usage_by_month",
		//	DataStruct:  &fastly.ServicesByRegionsUsage{},
		//	Description: "https://developer.fastly.com/reference/api/metrics-stats/historical-stats/#get-usage-month",
		//	PKColumns:   []string{},
		//	SkipFields:  []string{},
		//	Relations:   []string{},
		//},
		{
			TableName:   "stats_regions",
			DataStruct:  &struct{}{},
			Description: "https://developer.fastly.com/reference/api/metrics-stats/historical-stats/#get-regions",
			PKColumns:   []string{"name"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `setRegionName`,
				},
			},
		},
	}
	for _, r := range resources {
		r.Service = "stats"
	}
	return resources
}
