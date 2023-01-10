package stats

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StatsServices() *schema.Table {
	return &schema.Table{
		Name:        "fastly_stats_services",
		Description: `https://developer.fastly.com/reference/api/metrics-stats/historical-stats/`,
		Resolver:    fetchStatsServices,
		Multiplex:   client.ServiceRegionMultiplex,
		Transform:   transformers.TransformWithStruct(&models.StatsWrapper{}, transformers.WithSkipFields("MissHistogram")),
		Columns: []schema.Column{
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.UnixTimeResolver("StartTime"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("By"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
