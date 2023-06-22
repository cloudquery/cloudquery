package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"golang.org/x/sync/errgroup"
)

func (c Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.Message) error {
	filtered := schema.Tables{}
	for _, table := range c.tables {
		if !plugin.MatchesTable(table.Name, options.Tables, options.SkipTables) {
			continue
		}
		filtered = append(filtered, table)
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

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- message.Message) error {
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
			err := reverseTransform(rb.Field(i), values[i])
			if err != nil {
				return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
			}
		}
		res <- &message.Insert{Record: rb.NewRecord()}
	}
	return nil
}

func (c *Client) syncTables(ctx context.Context, tables schema.Tables, res chan<- message.Message) error {
	group, gctx := errgroup.WithContext(ctx)
	for _, table := range tables {
		if err := c.syncTable(gctx, table, res); err != nil {
			return err
		}
	}
	return group.Wait()
}
