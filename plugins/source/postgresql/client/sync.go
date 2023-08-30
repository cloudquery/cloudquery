package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	var err error
	var snapshotName string

	connPool, err := c.Conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	// this must be closed only at the end of the initial sync process otherwise the snapshot
	// used to sync the initial data will be released
	defer connPool.Release()
	conn := connPool.Conn().PgConn()

	filteredTables, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	for _, table := range filteredTables {
		res <- &message.SyncMigrateTable{
			Table: table,
		}
	}

	if c.pluginSpec.CDCId != "" {
		snapshotName, err = c.startCDC(ctx, filteredTables, conn)
		if err != nil {
			return err
		}
	}

	if c.pluginSpec.CDCId != "" && snapshotName == "" {
		c.logger.Info().Msg("cdc is enabled but replication slot already exists, skipping initial sync")
	} else {
		if err := c.syncTables(ctx, snapshotName, filteredTables, res); err != nil {
			return err
		}
	}

	if c.pluginSpec.CDCId == "" {
		return nil
	}

	if err := c.listenCDC(ctx, res); err != nil {
		return fmt.Errorf("failed to listen to cdc: %w", err)
	}
	return nil
}

func (c *Client) syncTables(ctx context.Context, snapshotName string, filteredTables schema.Tables, res chan<- message.SyncMessage) error {
	tx, err := c.Conn.BeginTx(ctx, pgx.TxOptions{
		// this transaction is needed for us to take a snapshot and we need to close it only at the end of the initial sync
		// https://www.postgresql.org/docs/current/transaction-iso.html
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if !errors.Is(err, pgx.ErrTxClosed) {
				c.logger.Error().Err(err).Msg("failed to rollback sync transaction")
			}
		}
	}()

	if snapshotName != "" {
		// if we use cdc we need to set the snapshot that was exported when we started the logical
		// replication stream
		if _, err := tx.Exec(ctx, "SET TRANSACTION SNAPSHOT '"+snapshotName+"'"); err != nil {
			return fmt.Errorf("failed to 'SET TRANSACTION SNAPSHOT %s': %w", snapshotName, err)
		}
	}

	for _, table := range filteredTables {
		if err := c.syncTable(ctx, tx, table, res); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit sync transaction: %w", err)
	}
	return nil
}

func (c *Client) syncTable(ctx context.Context, tx pgx.Tx, table *schema.Table, res chan<- message.SyncMessage) error {
	arrowSchema := table.ToArrowSchema()
	builder := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
	transformers := transformersForSchema(arrowSchema)

	colNames := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		colNames = append(colNames, pgx.Identifier{col.Name}.Sanitize())
	}
	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + pgx.Identifier{table.Name}.Sanitize()
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	rowsInRecord := 0
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}

		for i, value := range values {
			val, err := transformers[i](value)
			if err != nil {
				return err
			}

			s := scalar.NewScalar(arrowSchema.Field(i).Type)
			if err := s.Set(val); err != nil {
				return err
			}

			scalar.AppendToBuilder(builder.Field(i), s)
		}

		rowsInRecord++
		if rowsInRecord >= c.pluginSpec.RowsPerRecord {
			res <- &message.SyncInsert{Record: builder.NewRecord()} // NewRecord resets the builder for reuse
			rowsInRecord = 0
		}
	}

	record := builder.NewRecord()
	if record.NumRows() > 0 { // only send if there are some unsent rows
		res <- &message.SyncInsert{Record: record}
	}

	return nil
}

func stringForTime(t pgtype.Time, unit arrow.TimeUnit) string {
	extra := ""
	hour := t.Microseconds / 1e6 / 60 / 60
	minute := t.Microseconds / 1e6 / 60 % 60
	second := t.Microseconds / 1e6 % 60
	micros := t.Microseconds % 1e6
	switch unit {
	case arrow.Millisecond:
		extra = fmt.Sprintf(".%03d", (micros)/1e3)
	case arrow.Microsecond:
		extra = fmt.Sprintf(".%06d", micros)
	case arrow.Nanosecond:
		// postgres doesn't support nanosecond precision
		extra = fmt.Sprintf(".%06d", micros)
	}

	return fmt.Sprintf("%02d:%02d:%02d"+extra, hour, minute, second)
}
