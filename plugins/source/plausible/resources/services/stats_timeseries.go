package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/plausible/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StatsTimeseries() *schema.Table {
	return &schema.Table{
		Name:        "plausible_stats_timeseries",
		Description: "https://plausible.io/docs/stats-api#get-apiv1statstimeseries",
		Resolver:    fetchStatsTimeseries,
		Columns: []schema.Column{
			{
				Name:        "site_id",
				Description: "The site ID",
				Type:        schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
				Resolver: client.ResolveSiteID,
			},
			{
				Name:        "date",
				Description: "Date of the data point",
				Type:        schema.TypeTimestamp,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
				Resolver: ResolveDate,
			},
			{
				Name:        "visitors",
				Description: "The number of unique visitors",
				Type:        schema.TypeInt,
			},
			{
				Name:        "page_views",
				Description: "The number of pageview events",
				Type:        schema.TypeInt,
			},
			{
				Name:        "bounce_rate",
				Description: "Bounce rate percentage",
				Type:        schema.TypeInt,
			},
			{
				Name:        "visit_duration",
				Description: "Visit duration in seconds",
				Type:        schema.TypeInt,
			},
			{
				Name:        "visits",
				Description: "The number of visits/sessions",
				Type:        schema.TypeInt,
			},
		},
	}
}

type StatsTimeseriesResult struct {
	Date          string `json:"date"`
	Visitors      int    `json:"visitors"`
	PageViews     int    `json:"page_views"`
	BounceRate    int    `json:"bounce_rate"`
	VisitDuration int    `json:"visit_duration"`
	Visits        int    `json:"visits"`
}

type StatsTimeseriesResponse struct {
	Results []StatsTimeseriesResult `json:"results"`
}

func ResolveDate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	result := resource.Item.(StatsTimeseriesResult)
	t, err := time.Parse("2006-01-02", result.Date)
	if err != nil {
		return err
	}
	return resource.Set("date", t)
}

func fetchStatsTimeseries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	response := StatsTimeseriesResponse{}
	req := fmt.Sprintf("/api/v1/stats/timeseries?site_id=%s&metrics=%s&interval=%s&period=%s", cl.PluginSpec.SiteId, strings.Join(cl.PluginSpec.Metrics, ","), cl.PluginSpec.Interval, cl.PluginSpec.Period)
	if cl.PluginSpec.Filters != "" {
		req += fmt.Sprintf("&filters=%s", cl.PluginSpec.Filters)
	}

	if err := cl.Get(ctx, req, &response); err != nil {
		return err
	}
	res <- response.Results
	return nil
}
