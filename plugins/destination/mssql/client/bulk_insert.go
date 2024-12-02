package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (c *Client) bulkInsert(ctx context.Context, tx *sql.Tx, table *schema.Table, records []arrow.Record) error {
	stmt, err := tx.PrepareContext(ctx,
		mssql.CopyIn(queries.SanitizedTableName(c.spec.Schema, table),
			mssql.BulkOptions{
				KeepNulls:         true,
				KilobytesPerBatch: int(c.spec.BatchSizeBytes >> 10),
				RowsPerBatch:      int(c.spec.BatchSize),
				Tablock:           true,
			},
			table.Columns.Names()...,
		),
	)
	if err != nil {
		return err
	}

	rows, err := queries.GetRows(array.NewTableFromRecords(table.ToArrowSchema(), records))
	if err != nil {
		return err
	}
	for _, row := range rows {
		if _, err := stmt.ExecContext(ctx, row...); err != nil {
			return err
		}
	}

	// send bulkInsert
	_, err = stmt.ExecContext(ctx)
	return err
}
