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
		case *arrow.Int8Type, *arrow.Uint8Type, *arrow.Int16Type, *arrow.Uint16Type, *arrow.Int32Type, *arrow.Uint32Type, *arrow.Int64Type, *arrow.Uint64Type:
			var r sql.NullInt64
			results = append(results, &r)
		case *arrow.Float16Type, *arrow.Float32Type, *arrow.Float64Type:
			var r sql.NullFloat64
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
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int8Builder).Append(int8(v.Int64))
			}
		case *arrow.Int16Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int16Builder).Append(int16(v.Int64))
			}
		case *arrow.Int32Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int32Builder).Append(int32(v.Int64))
			}
		case *arrow.Int64Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Int64Builder).Append(v.Int64)
			}
		case *arrow.Uint8Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint8Builder).Append(uint8(v.Int64))
			}
		case *arrow.Uint16Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint16Builder).Append(uint16(v.Int64))
			}
		case *arrow.Uint32Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint32Builder).Append(uint32(v.Int64))
			}
		case *arrow.Uint64Type:
			v := val.(*sql.NullInt64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Uint64Builder).Append(uint64(v.Int64))
			}
		case *arrow.Float32Type:
			v := val.(*sql.NullFloat64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Float32Builder).Append(float32(val.(*sql.NullFloat64).Float64))
			}
		case *arrow.Float64Type:
			v := val.(*sql.NullFloat64)
			if !v.Valid {
				recordBuilder.Field(i).AppendNull()
			} else {
				recordBuilder.Field(i).(*array.Float64Builder).Append(val.(*sql.NullFloat64).Float64)
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
					panic(err)
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
