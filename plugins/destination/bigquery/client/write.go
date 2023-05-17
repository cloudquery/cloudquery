package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"google.golang.org/api/googleapi"
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

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	inserter := c.client.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Inserter()
	inserter.IgnoreUnknownValues = true
	inserter.SkipInvalidRows = false
	batch := make([]*item, 0)
	for _, rec := range records {
		for r := 0; r < int(rec.NumRows()); r++ {
			saver := &item{
				cols: make(map[string]bigquery.Value, len(table.Columns)),
			}
			for i, col := range rec.Columns() {
				if col.IsNull(r) {
					// save some bandwidth by not sending nil values
					continue
				}
				saver.cols[table.Columns[i].Name] = c.getValueForBigQuery(col, r)
			}
			batch = append(batch, saver)
		}
	}
	// flush final rows
	timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
	defer cancel()

	for err := inserter.Put(timeoutCtx, batch); err != nil; err = inserter.Put(timeoutCtx, batch) {
		// check if bigquery error is 404 (table does not exist yet), then wait a bit and retry until it does exist
		if e, ok := err.(*googleapi.Error); ok && e.Code == 404 {
			// retry
			c.logger.Info().Str("table", table.Name).Msg("Table does not exist yet, waiting for it to be created before retrying write")
			time.Sleep(1 * time.Second)
			continue
		}
		return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) getValueForBigQuery(col arrow.Array, i int) any {
	switch col.DataType().ID() {
	case arrow.MAP, arrow.STRUCT:
		v := col.GetOneForMarshal(i)
		b, _ := json.Marshal(v)
		return string(b)
	case arrow.INTERVAL_MONTH_DAY_NANO:
		return col.GetOneForMarshal(i).(arrow.MonthDayNanoInterval)
	case arrow.INTERVAL_DAY_TIME:
		return col.GetOneForMarshal(i).(arrow.DayTimeInterval)
	}
	return col.GetOneForMarshal(i)
}
