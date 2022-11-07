package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v4"
)

const (
	readSQL = "SELECT * FROM %s WHERE _cq_source_name = $1"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- *schema.DestinationResource) error {
	sql := fmt.Sprintf(readSQL, pgx.Identifier{table.Name}.Sanitize())
	rows, err := c.conn.Query(ctx, sql, sourceName)
	if err != nil {
		return err
	}
	transformer := plugins.DefaultReverseTransformer{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}
		cqTypes, err := transformer.ReverseTransformValues(table, values)
		if err != nil {
			return err
		}
		res <- &schema.DestinationResource{
			TableName: table.Name,
			Data:      cqTypes,
		}
	}
	rows.Close()
	return nil
}
