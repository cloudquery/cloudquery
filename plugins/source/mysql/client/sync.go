package client

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func (c Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	filtered, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	for _, table := range filtered {
		res <- &message.SyncMigrateTable{
			Table: table,
		}
	}
	return c.syncTables(ctx, filtered, res)
}

func (*Client) createResultsArray(table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		// We only support types that we create based on the schema, see SchemaType function
		switch col.Type.(type) {
		case *types.UUIDType, *arrow.BinaryType:
			var r *[]byte
			results = append(results, &r)
		case *arrow.BooleanType:
			var r *bool
			results = append(results, &r)
		case *arrow.Int8Type:
			var r *int8
			results = append(results, &r)
		case *arrow.Int16Type:
			var r *int16
			results = append(results, &r)
		case *arrow.Int32Type:
			var r *int32
			results = append(results, &r)
		case *arrow.Int64Type:
			var r *int64
			results = append(results, &r)
		case *arrow.Uint8Type:
			var r *uint8
			results = append(results, &r)
		case *arrow.Uint16Type:
			var r *uint16
			results = append(results, &r)
		case *arrow.Uint32Type:
			var r *uint32
			results = append(results, &r)
		case *arrow.Uint64Type:
			var r *uint64
			results = append(results, &r)
		case *arrow.Decimal128Type:
			var r *string
			results = append(results, &r)
		case *arrow.Float32Type:
			var r *float32
			results = append(results, &r)
		case *arrow.Float64Type:
			var r *float64
			results = append(results, &r)
		case *arrow.TimestampType:
			var r *time.Time
			results = append(results, &r)
		default:
			var r *string
			results = append(results, &r)
		}
	}
	return results
}

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- message.SyncMessage) error {
	colNames := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		colNames[i] = Identifier(col.Name)
	}
	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + Identifier(table.Name)
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		if err != nil {
			return err
		}

		arrowSchema := table.ToArrowSchema()
		rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
		for i := range values {
			// Gets the underlying value of the pointer
			val := reflect.ValueOf(values[i]).Elem().Interface()
			s := scalar.NewScalar(arrowSchema.Field(i).Type)
			if err := s.Set(val); err != nil {
				return err
			}
			scalar.AppendToBuilder(rb.Field(i), s)
		}
		res <- &message.SyncInsert{Record: rb.NewRecord()}
	}
	return nil
}

func (c *Client) syncTables(ctx context.Context, tables schema.Tables, res chan<- message.SyncMessage) error {
	for _, table := range tables {
		if err := c.syncTable(ctx, table, res); err != nil {
			return err
		}
	}
	return nil
}
