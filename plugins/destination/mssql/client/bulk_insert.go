package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (c *Client) bulkInsert(ctx context.Context, tx *sql.Tx, table *schema.Table, data [][]any) error {
	stmt, err := tx.PrepareContext(ctx,
		mssql.CopyIn(queries.SanitizedTableName(c.schemaName, table),
			mssql.BulkOptions{
				KeepNulls:         true,
				KilobytesPerBatch: c.spec.BatchSizeBytes >> 10,
				RowsPerBatch:      c.spec.BatchSize,
				Tablock:           true,
			},
			table.Columns.Names()...,
		),
	)
	if err != nil {
		return err
	}

	for _, d := range data {
		if _, err := stmt.ExecContext(ctx, d...); err != nil {
			return err
		}
	}

	// send bulkInsert
	_, err = stmt.ExecContext(ctx)
	return err
}
