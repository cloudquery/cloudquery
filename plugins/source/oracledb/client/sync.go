package client

import (
	"context"
	"fmt"
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
	"golang.org/x/sync/errgroup"
)

func (c Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
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
		for i, val := range values {
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
	group, gctx := errgroup.WithContext(ctx)
	group.SetLimit(c.concurrency)
	for _, table := range tables {
		t := table
		group.Go(func() error {
			return c.syncTable(gctx, t, res)
		})
	}
	return group.Wait()
}
