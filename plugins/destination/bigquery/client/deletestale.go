package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(table.Name)
		sb.WriteString(" where ")
		sb.WriteString(`"` + schema.CqSourceNameColumn.Name + `"`)
		sb.WriteString(" = @cq_source_name and \"")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString("\"::timestamp_tz < @cq_sync_time::timestamp_tz")
		sql := sb.String()
		q := c.client.Query(sql)
		q.Parameters = []bigquery.QueryParameter{
			{
				Name:  "cq_source_name",
				Value: source,
			},
			{
				Name:  "cq_sync_time",
				Value: syncTime,
			},
		}
		job, err := q.Run(ctx)
		if err != nil {
			return fmt.Errorf("failed to run query to delete stale entries from table %s: %w", table.Name, err)
		}
		js, err := job.Wait(ctx)
		if err != nil {
			return fmt.Errorf("failed to wait for job to delete stale entries from table %s: %w", table.Name, err)
		}
		if js.Err() != nil {
			return fmt.Errorf("job failed to delete stale entries from table %s: %w", table.Name, js.Err())
		}
	}
	return nil
}
