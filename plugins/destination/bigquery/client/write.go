package client

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	writeTimeout = 5 * time.Minute
)

type item struct {
	cols map[string]bigquery.Value
}

func (i *item) Save() (map[string]bigquery.Value, string, error) {
	// we're not doing de-dup at the moment
	return i.cols, bigquery.NoDedupeID, nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	inserter := c.client.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Inserter()
	inserter.IgnoreUnknownValues = true
	inserter.SkipInvalidRows = false
	batch := make([]*item, 0)
	for _, cols := range resources {
		saver := &item{
			cols: make(map[string]bigquery.Value, len(table.Columns)),
		}
		for i := range cols {
			if cols[i] == nil {
				// save some bandwidth by not sending nil values
				continue
			}
			saver.cols[table.Columns[i].Name] = cols[i]
		}
		batch = append(batch, saver)
	}
	// flush final rows
	timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
	defer cancel()
	err := inserter.Put(timeoutCtx, batch)
	if err != nil {
		return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
	}

	return nil
}
