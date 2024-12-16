package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const (
	writeTimeout = 5 * time.Minute
)

type item struct {
	Cols map[string]bigquery.Value `json:"cols"`
}

func (i *item) Save() (map[string]bigquery.Value, string, error) {
	// we're not doing de-dup at the moment
	return i.Cols, bigquery.NoDedupeID, nil
}

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, res); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}

func (c *Client) serializeBatchForError(batch []*item) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(batch)
	if err != nil {
		c.logger.Info().Err(err).Msg("Failed to serialize batch for error")
		return ""
	}
	const maxBufferSize = 1000
	if buffer.Len() > maxBufferSize {
		buffer.Truncate(maxBufferSize)
		buffer.WriteString("... (truncated)")
	}
	return buffer.String()
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	inserter := c.client.Dataset(c.spec.DatasetID).Table(name).Inserter()
	inserter.IgnoreUnknownValues = true
	inserter.SkipInvalidRows = false
	batch := make([]*item, 0)
	for _, msg := range msgs {
		rec := msg.Record
		for i := 0; i < int(rec.NumRows()); i++ {
			saver := &item{Cols: make(map[string]bigquery.Value, rec.NumCols())}
			for n, col := range rec.Columns() {
				if col.IsValid(i) {
					// save some bandwidth by not sending nil values
					saver.Cols[rec.ColumnName(n)] = getValueForBigQuery(col, i)
				}
			}
			batch = append(batch, saver)
		}
	}
	// flush final rows
	timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
	defer cancel()

	for err := inserter.Put(timeoutCtx, batch); err != nil; err = inserter.Put(timeoutCtx, batch) {
		// check if bigquery error is 404 (table does not exist yet), then wait a bit and retry until it does exist
		if isAPINotFoundError(err) {
			// retry
			c.logger.Info().Str("table", name).Msg("Table does not exist yet, waiting for it to be created before retrying write")
			time.Sleep(1 * time.Second)
			continue
		}
		if isEntityTooLargeError(err) {
			batchData := c.serializeBatchForError(batch)
			if batchData == "" {
				return fmt.Errorf("batch too big to be inserted into BigQuery table %s. See limitations here https://cloud.google.com/bigquery/quotas#streaming_inserts", name)
			}
			return fmt.Errorf("batch too big to be inserted into BigQuery table %s. See limitations here https://cloud.google.com/bigquery/quotas#streaming_inserts. Batch data: %s", name, batchData)
		}
		return fmt.Errorf("failed to put item into BigQuery table %s: %w", name, err)
	}

	return nil
}

func getValueForBigQuery(col arrow.Array, i int) any {
	switch v := col.(type) {
	case *array.Struct:
		m := map[string]bigquery.Value{}
		fields := v.DataType().(*arrow.StructType).Fields()
		for f, field := range fields {
			m[field.Name] = getValueForBigQuery(v.Field(f), i)
		}
		return m
	case *array.Map:
		v2 := col.GetOneForMarshal(i)
		b, _ := json.Marshal(v2)
		return string(b)
	case array.ListLike:
		col := col.(array.ListLike)
		from, to := col.ValueOffsets(i)
		slc := array.NewSlice(col.ListValues(), from, to)
		elems := make([]any, 0, slc.Len())
		for j := 0; j < slc.Len(); j++ {
			if slc.IsNull(j) {
				// LIMITATION: BigQuery does not support null values in repeated columns.
				// Therefore, these get stripped out here. In the future, perhaps we should support
				// an option to use JSON instead of repeated columns for users who need to preserve
				// the null values.
				continue
			}
			elems = append(elems, getValueForBigQuery(slc, j))
		}
		return elems
	case *array.MonthDayNanoInterval:
		return v.Value(i)
	case *array.DayTimeInterval:
		return v.Value(i)
	case *array.Duration:
		return v.Value(i)
	case *array.Timestamp:
		return v.Value(i).ToTime(v.DataType().(*arrow.TimestampType).Unit)
	case *types.JSONArray:
		return v.ValueStr(i)
	}
	return col.GetOneForMarshal(i)
}
