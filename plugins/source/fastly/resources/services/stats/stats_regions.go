package stats

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func StatsRegions() *schema.Table {
	return &schema.Table{
		Name:        "fastly_stats_regions",
		Description: `https://developer.fastly.com/reference/api/metrics-stats/historical-stats/#get-regions`,
		Resolver:    fetchStatsRegions,
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   setRegionName,
				PrimaryKey: true,
			},
		},
	}
}
