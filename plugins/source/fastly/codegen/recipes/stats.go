package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StatsResources() []*Resource {
	resources := []*Resource{
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
		{
			TableName:   "stats_services",
			DataStruct:  &models.StatsWrapper{},
			Description: "https://developer.fastly.com/reference/api/metrics-stats/historical-stats/",
			PKColumns:   []string{"service_id", "region", "start_time", "by"},
			Multiplex:   `client.ServiceRegionMultiplex`,
			SkipFields: []string{
				"StartTime",     // custom column resolver
				"MissHistogram", // decoding issues
			},
			UnwrapEmbeddedStructs: true,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "start_time",
					Type:     schema.TypeTimestamp,
					Resolver: `client.UnixTimeResolver("StartTime")`,
				},
			},
		},
	}
	for _, r := range resources {
		r.Service = "stats"
	}
	return resources
}
