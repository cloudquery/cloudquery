package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/goccy/go-json"
	"google.golang.org/api/iterator"
)

const (
	readSQL = "SELECT %s FROM `%s.%s.%s` WHERE `_cq_source_name` = @cq_source_name order by _cq_sync_time asc"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
	colSQL := "`" + strings.Join(table.Columns.Names(), "`, `") + "`"
	stmt := fmt.Sprintf(readSQL, colSQL, c.pluginSpec.ProjectID, c.pluginSpec.DatasetID, table.Name)
	q := c.client.Query(stmt)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "cq_source_name",
			Value: sourceName,
		},
	}
	q.Location = c.client.Location
	it, err := q.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to read table %s: %w", table.Name, err)
	}
	arrowSchema := table.ToArrowSchema()
	for {
		values := c.createResultsArray(arrowSchema)
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
		for i := range values {
			b, _ := json.Marshal(values[i])
			r := strings.NewReader(string(b))
			d := json.NewDecoder(r)
			err = rb.Field(i).UnmarshalOne(d)
			if err != nil {
				return fmt.Errorf("failed to unmarshal value: %w", err)
			}
		}
		res <- rb.NewRecord()
	}
	return nil
}

func (*Client) createResultsArray(sc *arrow.Schema) []bigquery.Value {
	results := make([]bigquery.Value, 0, len(sc.Fields()))
	for _, col := range sc.Fields() {
		dt := col.Type
		switch {
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Boolean):
			var r bool
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int8):
			var r int8
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int16):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int32):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int64):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint32):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint64):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float32):
			var r float64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float64):
			var r float64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.BinaryTypes.String):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(arrow.BinaryTypes.String)):
			var r []string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeString):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(arrow.BinaryTypes.LargeString)):
			var r []string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.BinaryTypes.Binary):
			var r []byte
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeBinary):
			var r []byte
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date32):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date64):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_s):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ms):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_us):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ns):
			var r *time.Time
			results = append(results, &r)
		case arrow.TypeEqual(dt, types.ExtensionTypes.Inet):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Inet)):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, types.ExtensionTypes.MAC):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.MAC)):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, types.ExtensionTypes.UUID):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.UUID)):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, types.ExtensionTypes.JSON):
			var r string
			results = append(results, &r)
		default:
			// use a string for unsupported types
			var r string
			results = append(results, &r)
		}
	}
	return results
}
