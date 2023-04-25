package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"google.golang.org/api/iterator"
)

const (
	readSQL = "SELECT %s FROM `%s.%s.%s` WHERE `_cq_source_name` = @cq_source_name order by _cq_sync_time asc"
)

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	colNames := make([]string, 0, len(arrowSchema.Fields()))
	for _, col := range arrowSchema.Fields() {
		colNames = append(colNames, col.Name)
	}
	colSQL := "`" + strings.Join(colNames, "`, `") + "`"
	stmt := fmt.Sprintf(readSQL, colSQL, c.pluginSpec.ProjectID, c.pluginSpec.DatasetID, tableName)
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
		return fmt.Errorf("failed to read table %s: %w", tableName, err)
	}
	for {
		values := c.createResultsArray(arrowSchema)
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", tableName, err)
		}
		rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
		for i := range values {
			c.appendValue(rb.Field(i), values[i])
		}
		res <- rb.NewRecord()
	}
	return nil
}

func (*Client) appendValue(b array.Builder, val bigquery.Value) {
	dt := b.Type()
	switch {
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Boolean):
		b.(*array.BooleanBuilder).Append(val.(bool))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int8):
		b.(*array.Int8Builder).Append(val.(int8))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int16):
		b.(*array.Int16Builder).Append(val.(int16))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int32):
		b.(*array.Int32Builder).Append(val.(int32))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int64):
		b.(*array.Int64Builder).Append(val.(int64))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
		b.(*array.Uint8Builder).Append(val.(uint8))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
		b.(*array.Uint16Builder).Append(val.(uint16))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint32):
		b.(*array.Uint32Builder).Append(val.(uint32))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint64):
		b.(*array.Uint64Builder).Append(val.(uint64))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float32):
		b.(*array.Float32Builder).Append(val.(float32))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float64):
		b.(*array.Float64Builder).Append(val.(float64))
	case arrow.TypeEqual(dt, arrow.BinaryTypes.String):
		b.(*array.StringBuilder).Append(val.(string))
	case arrow.TypeEqual(dt, arrow.ListOf(arrow.BinaryTypes.String)):
		b.(*array.ListBuilder).Append(true)
		vb := b.(*array.ListBuilder).ValueBuilder().(*array.StringBuilder)
		vb.AppendValues(val.([]string), nil)
	case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeString):
		b.(*array.LargeStringBuilder).Append(val.(string))
	case arrow.TypeEqual(dt, arrow.ListOf(arrow.BinaryTypes.LargeString)):
		b.(*array.ListBuilder).Append(true)
		vb := b.(*array.ListBuilder).ValueBuilder().(*array.LargeStringBuilder)
		vb.AppendValues(val.([]string), nil)
	case arrow.TypeEqual(dt, arrow.BinaryTypes.Binary):
		b.(*array.BinaryBuilder).Append(val.([]byte))
	case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeBinary):
		b.(*array.BinaryBuilder).Append(val.([]byte))
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date32):
		t := val.(*time.Time)
		b.(*array.Date32Builder).Append(arrow.Date32FromTime(*t))
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date64):
		t := val.(*time.Time)
		b.(*array.Date64Builder).Append(arrow.Date64FromTime(*t))
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_s):
		t := val.(*time.Time)
		arrowTimestamp, err := arrow.TimestampFromString(t.Format(time.RFC3339), arrow.Second)
		if err != nil {
			panic(err)
		}
		b.(*array.TimestampBuilder).Append(arrowTimestamp)
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ms):
		t := val.(*time.Time)
		arrowTimestamp, err := arrow.TimestampFromString(t.Format(time.RFC3339), arrow.Millisecond)
		if err != nil {
			panic(err)
		}
		b.(*array.TimestampBuilder).Append(arrowTimestamp)
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_us):
		t := val.(*time.Time)
		arrowTimestamp, err := arrow.TimestampFromString(t.Format(time.RFC3339), arrow.Microsecond)
		if err != nil {
			panic(err)
		}
		b.(*array.TimestampBuilder).Append(arrowTimestamp)
	case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ns):
		t := val.(*time.Time)
		arrowTimestamp, err := arrow.TimestampFromString(t.Format(time.RFC3339), arrow.Nanosecond)
		if err != nil {
			panic(err)
		}
		b.(*array.TimestampBuilder).Append(arrowTimestamp)
	case arrow.TypeEqual(dt, types.ExtensionTypes.Inet):
		b.(*types.InetBuilder).AppendValueFromString(val.(string))
	case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Inet)):
		b.(*array.ListBuilder).Append(true)
		vb := b.(*array.ListBuilder).ValueBuilder().(*types.InetBuilder)
		for _, s := range val.([]string) {
			vb.AppendValueFromString(s)
		}
	case arrow.TypeEqual(dt, types.ExtensionTypes.Mac):
		b.(*types.MacBuilder).AppendValueFromString(val.(string))
	case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Mac)):
		b.(*array.ListBuilder).Append(true)
		vb := b.(*array.ListBuilder).ValueBuilder().(*types.MacBuilder)
		for _, s := range val.([]string) {
			vb.AppendValueFromString(s)
		}
	case arrow.TypeEqual(dt, types.ExtensionTypes.UUID):
		b.(*types.UUIDBuilder).AppendValueFromString(val.(string))
	case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.UUID)):
		b.(*array.ListBuilder).Append(true)
		vb := b.(*array.ListBuilder).ValueBuilder().(*types.UUIDBuilder)
		for _, s := range val.([]string) {
			vb.AppendValueFromString(s)
		}
	case arrow.TypeEqual(dt, types.ExtensionTypes.JSON):
		b.(*types.JSONBuilder).AppendValueFromString(val.(string))
	default:
		panic(fmt.Sprintf("unsupported type for reading: %v", dt))
	}
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
			var r int16
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int32):
			var r int32
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int64):
			var r int64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
			var r uint8
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
			var r uint16
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint32):
			var r uint32
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint64):
			var r uint64
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float32):
			var r float32
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
		case arrow.TypeEqual(dt, types.ExtensionTypes.Mac):
			var r string
			results = append(results, &r)
		case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Mac)):
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
			panic(fmt.Sprintf("unsupported type for col %v: %v", col.Name, col.Type))
		}
	}
	return results
}
