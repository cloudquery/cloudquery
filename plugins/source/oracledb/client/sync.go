package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"golang.org/x/sync/errgroup"
)

func (c *Client) Sync(ctx context.Context, metrics *source.Metrics, res chan<- *schema.Resource) error {
	c.metrics = metrics
	for _, table := range c.Tables {
		if c.metrics.TableClient[table.Name] == nil {
			c.metrics.TableClient[table.Name] = make(map[string]*source.TableClientMetrics)
			c.metrics.TableClient[table.Name][c.ID()] = &source.TableClientMetrics{}
		}
	}

	return c.syncTables(ctx, res)
}

func (*Client) createResultsArray(table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		switch col.Type {
		case types.ExtensionTypes.UUID, arrow.BinaryTypes.Binary, arrow.BinaryTypes.LargeBinary:
			var r *[]byte
			results = append(results, &r)
		case arrow.FixedWidthTypes.Boolean:
			var r *bool
			results = append(results, &r)
		case arrow.PrimitiveTypes.Int8, arrow.PrimitiveTypes.Int16, arrow.PrimitiveTypes.Int32, arrow.PrimitiveTypes.Int64,
			arrow.PrimitiveTypes.Uint8, arrow.PrimitiveTypes.Uint16, arrow.PrimitiveTypes.Uint32, arrow.PrimitiveTypes.Uint64:
			var r *int
			results = append(results, &r)
		case arrow.FixedWidthTypes.Float16, arrow.FixedWidthTypes.Float16, arrow.PrimitiveTypes.Float32, arrow.PrimitiveTypes.Float64:
			var r *float64
			results = append(results, &r)
		case arrow.FixedWidthTypes.Timestamp_us:
			var r *time.Time
			results = append(results, &r)
		case types.ExtensionTypes.JSON:
			var r string
			results = append(results, &r)
		default:
			var r *string
			results = append(results, &r)
		}
	}
	return results
}

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- *schema.Resource) error {
	colNames := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		colNames[i] = Identifier(col.Name)
	}
	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + Identifier(table.Name)
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		c.metrics.TableClient[table.Name][c.ID()].Errors++
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		if err != nil {
			c.metrics.TableClient[table.Name][c.ID()].Errors++
			return err
		}
		resource, err := c.resourceFromValues(table.Name, values)
		if err != nil {
			c.metrics.TableClient[table.Name][c.ID()].Errors++
			return err
		}
		c.metrics.TableClient[table.Name][c.ID()].Resources++
		res <- resource
	}
	return nil
}

func (c *Client) syncTables(ctx context.Context, res chan<- *schema.Resource) error {
	group, gctx := errgroup.WithContext(ctx)
	group.SetLimit(int(c.Concurrency))
	for _, table := range c.Tables {
		if err := c.syncTable(gctx, table, res); err != nil {
			return err
		}
	}
	return group.Wait()
}

func (c *Client) resourceFromValues(tableName string, values []any) (*schema.Resource, error) {
	table := c.Tables.Get(tableName)
	resource := schema.NewResourceData(table, nil, values)
	for i, col := range table.Columns {
		if err := resource.Set(col.Name, values[i]); err != nil {
			return nil, err
		}
	}
	return resource, nil
}
