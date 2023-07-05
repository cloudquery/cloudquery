package stats

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "start_time",
				Type:       arrow.FixedWidthTypes.Timestamp_us,
				Resolver:   client.UnixTimeResolver("StartTime"),
				PrimaryKey: true,
			},
			{
				Name:       "service_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ServiceID"),
				PrimaryKey: true,
			},
			{
				Name:       "region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Region"),
				PrimaryKey: true,
			},
			{
				Name:       "by",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("By"),
				PrimaryKey: true,
			},
		},
	}
}
