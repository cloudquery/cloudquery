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
	for _, table := range c.Tables {
		if err := c.syncTable(ctx, table, res); err != nil {
			return err
		}
	}
	return nil
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
