package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

const concurrentDeletions = 10

// DeleteStale for BigQuery deletes stale records from previous syncs. However, since BigQuery uses a streaming buffer that
// stores records for up to 90 minutes and disallows deletions during this time, it is important that syncs using the delete-stale
// functionality be spaced at least 90 minutes apart. This shouldn't be a problem for most CloudQuery users, who typically
// run syncs once a day. If more frequent syncs are needed, consider batch-loading from a GCS bucket.
// See https://stackoverflow.com/questions/43085896/update-or-delete-tables-with-streaming-buffer-in-bigquery for more
// details and further reading.
func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	eg := errgroup.Group{}
	eg.SetLimit(concurrentDeletions)
	for _, table := range tables {
		table := table
		eg.Go(func() error {
			var sb strings.Builder
			sb.WriteString("delete from ")
			sb.WriteString("`" + c.datasetID + "." + table.Name + "`")
			sb.WriteString(" where ")
			sb.WriteString(schema.CqSourceNameColumn.Name)
			sb.WriteString(" = @cq_source_name and ")
			sb.WriteString(schema.CqSyncTimeColumn.Name)
			sb.WriteString(" < TIMESTAMP_SUB(TIMESTAMP(@cq_sync_time), INTERVAL 90 MINUTE)")
			sql := sb.String()
			q := c.client.Query(sql)
			q.Parameters = []bigquery.QueryParameter{
				{
					Name:  "cq_source_name",
					Value: source,
				},
				{
					Name:  "cq_sync_time",
					Value: syncTime.Format(time.RFC3339),
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
			return nil
		})
	}
	return eg.Wait()
}
