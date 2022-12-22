package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func CrowdScore() []*Resource {
	resources := []*Resource{
		{
			Service:    "incidents",
			SubService: "crowdscore",
			Struct:     &models.DomainEnvironmentScore{},
			PKColumns:  []string{"id"},
			SkipFields: []string{"Timestamp"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "timestamp",
					Type:     schema.TypeTimestamp,
					Resolver: `schema.PathResolver("Timestamp")`,
				},
			},
		},
	}
	return resources
}
