package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v4"
)

const (
	readSQL = "SELECT * FROM %s WHERE _cq_source_name = $1"
)


func (c *Client) Read(ctx context.Context, tables schema.Tables, sourceName string, res chan<- *schema.DestinationResource) error {
	for _, table := range tables {
		sql := fmt.Sprintf(readSQL, pgx.Identifier{table.Name}.Sanitize())
		rows, err := c.conn.Query(ctx, sql, sourceName)
		if err != nil {
			return err
		}
		for rows.Next() {
			rows.Values()
			values, err := rows.Values()
			if err != nil {
				return err
			}
			for _, v := range values {
				fmt.Printf("%T\n", v)
			}
			fmt.Println(values)
			cqTypes, err := schema.CQTypesFromValues(table, values)
			if err != nil {
				return err
			}
			res <- &schema.DestinationResource{
				TableName: table.Name,
				Data: cqTypes,
			}
		}
		rows.Close()
	}
	return nil
}
