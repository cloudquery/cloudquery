package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
)

func (c *Client) bulkInsert(ctx context.Context, tx *sql.Tx, sc *arrow.Schema, records []arrow.Record) error {
	return fmt.Errorf("unimplemented")
	//stmt, err := tx.PrepareContext(ctx,
	//	mssql.CopyIn(queries.SanitizedTableName(c.schemaName, sc),
	//		mssql.BulkOptions{
	//			KeepNulls:         true,
	//			KilobytesPerBatch: c.spec.BatchSizeBytes >> 10,
	//			RowsPerBatch:      c.spec.BatchSize,
	//			Tablock:           true,
	//		},
	//		queries.ColumnNames(sc)...,
	//	),
	//)
	//if err != nil {
	//	return err
	//}
	//
	//for _, record := range records {
	//	if _, err := stmt.ExecContext(ctx, d...); err != nil {
	//		return err
	//	}
	//}
	//
	//// send bulkInsert
	//_, err = stmt.ExecContext(ctx)
	//return err
}
