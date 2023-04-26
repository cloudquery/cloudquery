package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	query, params := queries.Read(c.schemaName, sourceName, table)

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, query, params...)
		if err != nil {
			return err
		}

		return processRows(rows, func(row *sql.Rows) error {
			// We consider only the current schema from table
			resource := rowArr(table)
			if err := row.Scan(resource...); err != nil {
				return err
			}

			res <- fixUUID(table, resource)

			return nil
		})
	})
}

func rowArr(table *schema.Table) []any {
	res := make([]any, len(table.Columns))
	for i := range res {
		res[i] = new(any)
	}
	return res
}

func unwrap(arr []any) {
	for i, v := range arr {
		arr[i] = *v.(*any)
	}
}

func fixUUID(table *schema.Table, resource []any) []any {
	unwrap(resource)
	for i, v := range resource {
		if table.Columns[i].Type == schema.TypeUUID {
			uid := mssql.UniqueIdentifier{}
			_ = uid.Scan(v)
			resource[i] = uid
		}
	}
	return resource
}
