package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
	sc := table.ToArrowSchema()
	query, params := queries.Read(c.schemaName, sourceName, table)
	l := len(sc.Fields())

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, query, params...)
		if err != nil {
			return err
		}

		return processRows(rows, func(row *sql.Rows) error {
			// We consider only the current schema from table
			resource := make([]any, l)
			if err := row.Scan(wrap(resource)...); err != nil {
				return err
			}

			record, err := queries.Record(sc, resource)
			if err != nil {
				return err
			}

			res <- record
			return nil
		})
	})
}

func wrap(arr []any) []any {
	res := make([]any, len(arr))
	for i := range res {
		res[i] = &arr[i]
	}
	return res
}
