package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
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
	switch v := col.(type) {
	case *array.Struct:
		m := map[string]bigquery.Value{}
		fields := v.DataType().(*arrow.StructType).Fields()
		for f, field := range fields {
			m[field.Name] = c.getValueForBigQuery(v.Field(f), i)
		}
		return m
	case *array.Map:
		v2 := col.GetOneForMarshal(i)
		b, _ := json.Marshal(v2)
		return string(b)
	case array.ListLike:
		arr := col.(array.ListLike)
		elems := make([]any, 0, arr.Len())
		for j := 0; j < arr.Len(); j++ {
			if arr.IsNull(j) {
				continue
			}
			from, to := arr.ValueOffsets(j)
			slc := array.NewSlice(arr.ListValues(), from, to)
			for k := 0; k < int(to-from); k++ {
				if slc.IsNull(k) {
					// LIMITATION: BigQuery does not support null values in repeated columns.
					// Therefore, these get stripped out here. In the future, perhaps we should support
					// an option to use JSON instead of repeated columns for users who need to preserve
					// the null values.
					continue
				}
				elems = append(elems, c.getValueForBigQuery(slc, k))
			}
		}
		return elems
	case *array.MonthDayNanoInterval:
		return v.Value(i)
	case *array.DayTimeInterval:
		return v.Value(i)
	case *array.Duration:
		return v.Value(i)
	case *array.Timestamp:
		unit := v.DataType().(*arrow.TimestampType).Unit
		switch unit {
		case arrow.Nanosecond:
			t := v.Value(i).ToTime(arrow.Nanosecond)
			format := "2006-01-02 15:04:05.999999"
			return TimestampNanoseconds{
				Timestamp:   t.Format(format),
				Nanoseconds: t.Nanosecond() % 1000,
			}
		default:
			return v.GetOneForMarshal(i)
		}
	}
	return col.GetOneForMarshal(i)
}
