package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
)

func (c *Client) Read(ctx context.Context, sc *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	query, params := queries.Read(c.schemaName, sourceName, sc)
	l := len(sc.Fields())

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, query, params...)
		if err != nil {
			return err
		}

		return processRows(rows, func(row *sql.Rows) error {
			// We consider only the current schema from table
			resource := rowArr(l)
			if err := row.Scan(resource...); err != nil {
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

func rowArr(l int) []any {
	res := make([]any, l)
	for i := range res {
		res[i] = new(any)
	}
	return res
}
