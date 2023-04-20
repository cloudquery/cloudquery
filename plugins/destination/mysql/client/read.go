package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func (*Client) createResultsArray(table *arrow.Schema) []any {
	results := make([]any, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		switch col.Type.(type) {
		case *arrow.BooleanType:
			var r sql.NullBool
			results = append(results, &r)
		case *arrow.Int8Type:
			var r *int8
			results = append(results, &r)
		case *arrow.Uint8Type:
			var r *uint8
			results = append(results, &r)
		case *arrow.Int16Type:
			var r *int16
			results = append(results, &r)
		case *arrow.Uint16Type:
			var r *uint16
			results = append(results, &r)
		case *arrow.Int32Type:
			var r *int32
			results = append(results, &r)
		case *arrow.Uint32Type:
			var r *uint32
			results = append(results, &r)
		case *arrow.Int64Type:
			var r *int64
			results = append(results, &r)
		case *arrow.Uint64Type:
			var r *uint64
			results = append(results, &r)
		case *arrow.Float32Type:
			var r *float32
			results = append(results, &r)
		case *arrow.Float64Type:
			var r *float64
			results = append(results, &r)
		case *arrow.StringType, *arrow.LargeStringType:
			var r sql.NullString
			results = append(results, &r)
		case *arrow.BinaryType, *arrow.LargeBinaryType:
			var r []byte
			results = append(results, &r)
		case *types.UUIDType:
			var r []byte
			results = append(results, &r)
		case *arrow.TimestampType:
			var r *time.Time
			results = append(results, &r)
		default:
			var r sql.NullString
			results = append(results, &r)
		}
	}
	return results
}

func reverseTransform(table *arrow.Schema, values []any) (arrow.Record, error) {
	recordBuilder := array.NewRecordBuilder(memory.DefaultAllocator, table)
	defer recordBuilder.Release()
	for i, val := range values {
		switch table.Field(i).Type.(type) {
		case *arrow.BooleanType:
			if val.(*sql.NullBool).Valid {
				recordBuilder.Field(i).(*array.BooleanBuilder).Append(val.(*sql.NullBool).Bool)
			} else {
				recordBuilder.Field(i).(*array.BooleanBuilder).AppendNull()
			}
		case *arrow.Int8Type:
			v := val.(**int8)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int8Builder).Append(**v)
			}
		case *arrow.Int16Type:
			v := val.(**int16)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int16Builder).Append(**v)
			}
		case *arrow.Int32Type:
			v := val.(**int32)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int32Builder).Append(**v)
			}
		case *arrow.Int64Type:
			v := val.(**int64)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int64Builder).Append(**v)
			}
		case *arrow.Uint8Type:
			v := val.(**uint8)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint8Builder).Append(**v)
			}
		case *arrow.Uint16Type:
			v := val.(**uint16)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint16Builder).Append(**v)
			}
		case *arrow.Uint32Type:
			v := val.(**uint32)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint32Builder).Append(**v)
			}
		case *arrow.Uint64Type:
			v := val.(**uint64)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint64Builder).Append(**v)
			}
		case *arrow.Float32Type:
			v := val.(**float32)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Float32Builder).Append(**v)
			}
		case *arrow.Float64Type:
			v := val.(**float64)
			if *v == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Float64Builder).Append(**v)
			}
		case *arrow.StringType:
			v := val.(*sql.NullString)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.StringBuilder).Append(val.(*sql.NullString).String)
			}
		case *arrow.LargeStringType:
			v := val.(*sql.NullString)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.LargeStringBuilder).Append(val.(*sql.NullString).String)
			}
		case *arrow.BinaryType:
			if *val.(*[]byte) == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.BinaryBuilder).Append(*val.(*[]byte))
			}
		case *arrow.TimestampType:
			asTime := val.(**time.Time)
			if *asTime == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.TimestampBuilder).Append(arrow.Timestamp((*asTime).UnixMicro()))
			}
		case *types.UUIDType:
			if *val.(*[]byte) == nil {
				recordBuilder.Field(i).AppendNull()
			} else {
				asUUID, err := uuid.FromBytes(*val.(*[]byte))
				if err != nil {
					return nil, err
				}
				recordBuilder.Field(i).(*types.UUIDBuilder).Append(asUUID)
			}
		default:
			v := val.(*sql.NullString)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				if err := recordBuilder.Field(i).AppendValueFromString(val.(*sql.NullString).String); err != nil {
					return nil, fmt.Errorf("failed to AppendValueFromString %s. field: %v. name: %s err: %w", *val.(*string), recordBuilder.Field(i).Type(), table.Fields()[i].Name, err)
				}
			}
		}
	}
	rec := recordBuilder.NewRecord()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	builder := strings.Builder{}
	builder.WriteString("SELECT")
	fields := table.Fields()
	for i, col := range fields {
		builder.WriteString(" " + identifier(col.Name))
		if i != len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	tableName := schema.TableName(table)
	builder.WriteString("FROM " + identifier(tableName) + " WHERE _cq_source_name = ? ORDER BY _cq_sync_time ASC")
	rows, err := c.db.QueryContext(ctx, builder.String(), sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
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
	return nil
}
