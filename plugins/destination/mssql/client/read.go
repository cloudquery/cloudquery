package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	query := queries.Read(c.spec.Schema, table)
	sc := table.ToArrowSchema()

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, query)
		if err != nil {
			return err
		}

		return processRows(rows, func(row *sql.Rows) error {
			// We consider only the current schema from table
			resource := make([]any, len(sc.Fields()))
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
