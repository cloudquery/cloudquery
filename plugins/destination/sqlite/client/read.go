package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	readSQL = `SELECT %s FROM %s`
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
			var r sql.NullInt64
			results = append(results, &r)
		case arrow.FLOAT16, arrow.FLOAT32, arrow.FLOAT64:
			var r sql.NullFloat64
			results = append(results, &r)
		default:
			var r sql.NullString
			results = append(results, &r)
		}
	}
	return results
}

func reverseTransform(sc *arrow.Schema, values []any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, val := range values {
		switch sc.Field(i).Type.ID() {
		case arrow.BOOL:
			if val.(*sql.NullBool).Valid {
				bldr.Field(i).(*array.BooleanBuilder).Append(val.(*sql.NullBool).Bool)
			} else {
				bldr.Field(i).(*array.BooleanBuilder).AppendNull()
			}
		case arrow.INT8:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Int8Builder).Append(int8(v.Int64))
			}
		case arrow.INT16:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Int16Builder).Append(int16(v.Int64))
			}
		case arrow.INT32:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Int32Builder).Append(int32(v.Int64))
			}
		case arrow.INT64:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Int64Builder).Append(v.Int64)
			}
		case arrow.UINT8:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Uint8Builder).Append(uint8(v.Int64))
			}
		case arrow.UINT16:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Uint16Builder).Append(uint16(v.Int64))
			}
		case arrow.UINT32:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Uint32Builder).Append(uint32(v.Int64))
			}
		case arrow.UINT64:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Uint64Builder).Append(uint64(v.Int64))
			}
		case arrow.FLOAT32:
			v := val.(*sql.NullFloat64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Float32Builder).Append(float32(val.(*sql.NullFloat64).Float64))
			}
		case arrow.FLOAT64:
			v := val.(*sql.NullFloat64)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.Float64Builder).Append(val.(*sql.NullFloat64).Float64)
			}
		case arrow.STRING:
			v := val.(*sql.NullString)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.StringBuilder).Append(val.(*sql.NullString).String)
			}
		case arrow.BINARY, arrow.LARGE_BINARY:
			if *val.(*[]byte) == nil {
				bldr.Field(i).AppendNull()
			} else {
				bldr.Field(i).(*array.BinaryBuilder).Append(*val.(*[]byte))
			}
		default:
			v := val.(*sql.NullString)
			if !v.Valid {
				bldr.Field(i).AppendNull()
			} else {
				if err := bldr.Field(i).AppendValueFromString(val.(*sql.NullString).String); err != nil {
					return nil, fmt.Errorf("failed to AppendValueFromString %s. field: %v. name: %s err: %w", *val.(*string), bldr.Field(i).Type(), sc.Fields()[i].Name, err)
				}
			}
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	colNames := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		colNames[i] = identifier(col.Name)
	}
	cols := strings.Join(colNames, ", ")
	rows, err := c.db.Query(fmt.Sprintf(readSQL, cols, identifier(table.Name)))
	if err != nil {
		return err
	}
	arrowSchema := table.ToArrowSchema()
	for rows.Next() {
		values := c.createResultsArray(arrowSchema)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		record, err := reverseTransform(arrowSchema, values)
		if err != nil {
			return err
		}
		res <- record
	}
	rows.Close()
	return nil
}
