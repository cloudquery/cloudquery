package client

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/netip"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/goccy/go-json"
	"golang.org/x/exp/slices"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	readSQL = "SELECT %s FROM %s"
)

// Sync reads all records from the table and sends them to the channel
func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- plugin.Message) error {
	tables, err := c.listTables(ctx, options.Tables, options.SkipTables)
	if err != nil {
		return fmt.Errorf("failed to list tables: %w", err)
	}
	for _, table := range tables {
		if !slices.Contains(options.Tables, table.Name) {
			// TODO(v4): handle wildcards and skip tables
			continue
		}
		colNames := make([]string, 0, len(table.Columns))
		for _, col := range table.Columns {
			colNames = append(colNames, pgx.Identifier{col.Name}.Sanitize())
		}
		cols := strings.Join(colNames, ",")
		tableName := table.Name
		sql := fmt.Sprintf(readSQL, cols, pgx.Identifier{tableName}.Sanitize())
		rows, err := c.conn.Query(ctx, sql)
		if err != nil {
			return fmt.Errorf("failed to execute SQL read query: %w", err)
		}
		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}
			rec, err := c.reverseTransformer(table, values)
			if err != nil {
				return err
			}
			res <- &plugin.MessageInsert{
				Record: rec,
				Upsert: false,
			}
		}
		rows.Close()
	}
	return nil
}

// nolint: dupl
func (c *Client) reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}

	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		b.Append(val.(bool))
	case *array.Int8Builder:
		// pgx always return int16 for int8
		b.Append(int8(val.(int16)))
	case *array.Int16Builder:
		b.Append(val.(int16))
	case *array.Int32Builder:
		b.Append(val.(int32))
	case *array.Int64Builder:
		b.Append(val.(int64))
	case *array.Uint8Builder:
		b.Append(uint8(val.(int16)))
	case *array.Uint16Builder:
		b.Append(uint16(val.(int32)))
	case *array.Uint32Builder:
		b.Append(uint32(val.(int64)))
	case *array.Uint64Builder:
		v := val.(pgtype.Numeric)
		b.Append(v.Int.Uint64())
	case *array.Float32Builder:
		b.Append(val.(float32))
	case *array.Float64Builder:
		b.Append(val.(float64))
	case *array.StringBuilder:
		va, ok := val.(string)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T and column %s", val, bldr, f.Name)
		}
		b.Append(va)
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		b.Append(val.([]byte))
	case *array.TimestampBuilder:
		switch b.Type().(*arrow.TimestampType).Unit {
		case arrow.Second:
			b.Append(arrow.Timestamp(val.(time.Time).Unix()))
		case arrow.Millisecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixMilli()))
		case arrow.Microsecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixMicro()))
		case arrow.Nanosecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixNano()))
		default:
			return fmt.Errorf("unsupported timestamp unit %s", f.Type.(*arrow.TimestampType).Unit)
		}
	case *types.UUIDBuilder:
		va, ok := val.([16]byte)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T", val, bldr)
		}
		u, err := uuid.FromBytes(va[:])
		if err != nil {
			return err
		}
		b.Append(u)
	case *types.JSONBuilder:
		b.Append(val)
	case *array.StructBuilder:
		structBytes, err := json.Marshal(val)
		if err != nil {
			return fmt.Errorf("failed to marshal struct: %w", err)
		}
		dec := json.NewDecoder(bytes.NewReader(structBytes))
		if err := b.UnmarshalOne(dec); err != nil {
			return fmt.Errorf("failed to unmarshal struct: %w", err)
		}
	case *types.InetBuilder:
		if v, ok := val.(netip.Prefix); ok {
			_, ipnet, err := net.ParseCIDR(v.String())
			if err != nil {
				return err
			}
			b.Append(ipnet)
			return nil
		}
		b.Append(val.(*net.IPNet))
	case *types.MACBuilder:
		if c.pgType == pgTypePostgreSQL {
			b.Append(val.(net.HardwareAddr))
		} else {
			hardwareAddr, err := net.ParseMAC(val.(string))
			if err != nil {
				return err
			}
			b.Append(hardwareAddr)
		}
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		for _, v := range val.([]any) {
			if err := c.reverseTransform(f, valBuilder, v); err != nil {
				return err
			}
		}
	default:
		v, ok := val.(string)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T", val, bldr)
		}
		if err := bldr.AppendValueFromString(v); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) reverseTransformer(table *schema.Table, values []any) (arrow.Record, error) {
	sc := table.ToArrowSchema()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if c.pgType == pgTypePostgreSQL {
			if err := c.reverseTransform(f, bldr.Field(i), values[i]); err != nil {
				return nil, err
			}
		} else {
			if err := c.reverseTransformCockroach(f, bldr.Field(i), values[i]); err != nil {
				return nil, err
			}
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}
