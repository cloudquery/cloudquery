package client

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/iterator"
)

const (
	readSQL = "SELECT * FROM `%s.%s.%s` WHERE `_cq_source_name` = @cq_source_name"
)

func (*Client) createResultsArray(table *schema.Table) []bigquery.Value {
	results := make([]bigquery.Value, 0, len(table.Columns))
	for _, col := range table.Columns {
		switch col.Type {
		case schema.TypeBool:
			var r bool
			results = append(results, &r)
		case schema.TypeInt:
			var r int
			results = append(results, &r)
		case schema.TypeFloat:
			var r float64
			results = append(results, &r)
		case schema.TypeUUID:
			var r string
			results = append(results, &r)
		case schema.TypeString:
			var r string
			results = append(results, &r)
		case schema.TypeByteArray:
			var r sql.RawBytes
			results = append(results, &r)
		case schema.TypeStringArray:
			var r []string
			results = append(results, &r)
		case schema.TypeTimestamp:
			var r *time.Time
			results = append(results, &r)
		case schema.TypeJSON:
			var r string
			results = append(results, &r)
		case schema.TypeUUIDArray:
			var r []string
			results = append(results, &r)
		case schema.TypeCIDR:
			var r string
			results = append(results, &r)
		case schema.TypeCIDRArray:
			var r []string
			results = append(results, &r)
		case schema.TypeMacAddr:
			var r string
			results = append(results, &r)
		case schema.TypeMacAddrArray:
			var r []string
			results = append(results, &r)
		case schema.TypeInet:
			var r string
			results = append(results, &r)
		case schema.TypeInetArray:
			var r []string
			results = append(results, &r)
		case schema.TypeIntArray:
			var r []int64
			results = append(results, &r)
		default:
			panic(fmt.Sprintf("unsupported type for col %v: %v", col.Name, col.Type))
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	stmt := fmt.Sprintf(readSQL, c.pluginSpec.ProjectID, c.pluginSpec.DatasetID, table.Name)
	client, err := c.bqClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	q := client.Query(stmt)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "cq_source_name",
			Value: sourceName,
		},
	}
	it, err := q.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to read table %s: %w", table.Name, err)
	}
	values := c.createResultsArray(table)
	v := make([]interface{}, len(values))
	for {
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		for i := range values {
			v[i] = values[i]
		}
		res <- v
	}
	return nil
}
