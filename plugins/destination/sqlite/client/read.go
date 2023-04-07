package client

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/goccy/go-json"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	readSQL = `SELECT %s FROM "%s" WHERE _cq_source_name = $1 order by _cq_sync_time asc`
)

func (*Client) createResultsArray(table *arrow.Schema) []any {
	results := make([]any, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		switch col.Type.ID() {
		case arrow.BOOL:
			var r sql.NullBool
			results = append(results, &r)
		case arrow.BINARY, arrow.LARGE_BINARY:
			var r []byte
			results = append(results, &r)
		case arrow.INT8, arrow.INT16, arrow.INT32, arrow.INT64, arrow.UINT8, arrow.UINT16, arrow.UINT32, arrow.UINT64:
			var r int
			results = append(results, &r)
		case arrow.FLOAT16, arrow.FLOAT32, arrow.FLOAT64:
			var r float64
			results = append(results, &r)
		default:
			var r string
			results = append(results, &r)
		}
	}
	return results
}


func reverseTransform(sc *arrow.Schema, values []any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	defer bldr.Release()
	for i, val := range values {
		switch sc.Field(i).Type.ID() {
		case arrow.BOOL:
			if val.(*sql.NullBool).Valid {
				bldr.Field(i).(*array.BooleanBuilder).Append(val.(*sql.NullBool).Bool)
			} else {
				bldr.Field(i).(*array.BooleanBuilder).AppendNull()
			}
			// bldr.Field(i).(*array.BooleanBuilder).Append(*val.(*bool))
		case arrow.INT8:
			bldr.Field(i).(*array.Int8Builder).Append(int8(*val.(*int)))
		case arrow.INT16:
			bldr.Field(i).(*array.Int16Builder).Append(int16(*val.(*int)))
		case arrow.INT32:
			bldr.Field(i).(*array.Int32Builder).Append(int32(*val.(*int)))
		case arrow.INT64:
			bldr.Field(i).(*array.Int64Builder).Append(int64(*val.(*int)))
		case arrow.UINT8:
			bldr.Field(i).(*array.Uint8Builder).Append(uint8(*val.(*int)))
		case arrow.UINT16:
			bldr.Field(i).(*array.Uint16Builder).Append(uint16(*val.(*int)))
		case arrow.UINT32:
			bldr.Field(i).(*array.Uint32Builder).Append(uint32(*val.(*int)))
		case arrow.UINT64:
			bldr.Field(i).(*array.Uint64Builder).Append(uint64(*val.(*int)))
		case arrow.FLOAT32:
			bldr.Field(i).(*array.Float32Builder).Append(float32(*val.(*float64)))
		case arrow.FLOAT64:
			bldr.Field(i).(*array.Float64Builder).Append(*val.(*float64))
		case arrow.STRING:
			bldr.Field(i).(*array.StringBuilder).Append(*val.(*string))
		case arrow.BINARY:
			bldr.Field(i).(*array.BinaryBuilder).Append(*val.(*[]byte))
		case arrow.DATE32, arrow.DATE64,
		arrow.TIMESTAMP,
		arrow.TIME32, arrow.TIME64,
		arrow.INTERVAL_DAY_TIME,
		arrow.DECIMAL128, arrow.DECIMAL256:
			dec := json.NewDecoder(bytes.NewReader([]byte(`"` + *val.(*string) + `"`)))
			if err := bldr.Field(i).UnmarshalOne(dec); err != nil {
				return nil, fmt.Errorf("failed to unmarshal %s. field: %v. err: %w", *val.(*string), bldr.Field(i).Type(), err)
			}
		default:
			dec := json.NewDecoder(bytes.NewReader([]byte(*val.(*string) )))
			if err := bldr.Field(i).UnmarshalOne(dec); err != nil {
				return nil, fmt.Errorf("failed to unmarshal %s. field: %v. err: %w", *val.(*string), bldr.Field(i).Type(), err)
			}
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	colNames := make([]string, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		colNames = append(colNames, `"`+col.Name+`"`)
	}
	cols := strings.Join(colNames, ", ")
	tableName := schema.TableName(table)
	rows, err := c.db.Query(fmt.Sprintf(readSQL, cols, tableName), sourceName)
	if err != nil {
		return err
	}
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", tableName, err)
		}
		record, err := reverseTransform(table, values)
		if err != nil {
			return err
		}
		res <- record
	}
	rows.Close()
	return nil
}
