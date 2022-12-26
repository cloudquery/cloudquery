package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

type statsResponse struct {
	Data    []*models.StatsWrapper `mapstructure:"data"`
	Message string                 `mapstructure:"msg"`
	Meta    map[string]string      `mapstructure:"meta"`
	Status  string                 `mapstructure:"status"`
}

func fetchStatsServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	startTime := *c.Service.CreatedAt
	startTime = startTime.Truncate(time.Hour)
	endTime := time.Now().Add(-5 * time.Minute).Truncate(time.Hour) // avoid storing incomplete data

	// read 60 days of hourly data at a time
	for startTime.Before(endTime) {
		var resp statsResponse
		err := c.Fastly.GetStatsJSON(&fastly.GetStatsInput{
			Service: c.Service.ID,
			Region:  c.Region,
			// We only support hourly stats right now.
			// Please raise a ticket if you need stats at minutely or daily granularity:
			// https://github.com/cloudquery/cloudquery/issues
			By:   "hour",
			From: fmt.Sprintf("%d", startTime.Unix()),
			To:   fmt.Sprintf("%d", startTime.Add(30*24*time.Hour).Unix()),
		}, &resp)
		if err != nil {
			return err
		}
		for _, stat := range resp.Data {
			stat.By = "hour"
		}
		res <- resp.Data
		startTime = startTime.Add(60 * 24 * time.Hour)
	}
	return nil
}
