package client

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c *Client) Sync(ctx context.Context, metrics *source.Metrics, res chan<- *schema.Resource) error {
	// var conn *pgconn.PgConn
	var err error
	var snapshotName string
	c.metrics = metrics
	for _, table := range c.Tables {
		if c.metrics.TableClient[table.Name] == nil {
			c.metrics.TableClient[table.Name] = make(map[string]*source.TableClientMetrics)
			c.metrics.TableClient[table.Name][c.ID()] = &source.TableClientMetrics{}
		}
	}

	connPool, err := c.Conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	// this must be closed only at the end of the initial sync process otherwise the snapshot
	// used to sync the initial data will be released
	defer connPool.Release()
	conn := connPool.Conn().PgConn()

	if c.pluginSpec.CDC {
		snapshotName, err = c.startCDC(ctx, conn)
		if err != nil {
			return err
		}
	}

	if c.pluginSpec.CDC && snapshotName == "" {
		c.logger.Info().Msg("cdc is enabled but replication slot already exists, skipping initial sync")
	} else {
		if err := c.syncTables(ctx, snapshotName, res); err != nil {
			return err
		}
	}

	if !c.pluginSpec.CDC {
		return nil
	}

	if err := c.listenCDC(ctx, res); err != nil {
		return fmt.Errorf("failed to listen to cdc: %w", err)
	}
	return nil
}

func (c *Client) syncTable(ctx context.Context, tx pgx.Tx, table *schema.Table, res chan<- *schema.Resource) error {
	colNames := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		colNames[i] = pgx.Identifier{col.Name}.Sanitize()
	}
	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + pgx.Identifier{table.Name}.Sanitize()
	rows, err := tx.Query(ctx, query)
	if err != nil {
		c.metrics.TableClient[table.Name][c.ID()].Errors++
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			c.metrics.TableClient[table.Name][c.ID()].Errors++
			return err
		}
		resource, err := c.resourceFromValues(table.Name, values)
		if err != nil {
			c.metrics.TableClient[table.Name][c.ID()].Errors++
			return err
		}
		c.metrics.TableClient[table.Name][c.ID()].Resources++
		res <- resource
	}
	return nil
}

func (c *Client) syncTables(ctx context.Context, snapshotName string, res chan<- *schema.Resource) error {
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

	for _, table := range c.Tables {
		if err := c.syncTable(ctx, tx, table, res); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit sync transaction: %w", err)
	}
	return nil
}

func (c *Client) resourceFromValues(tableName string, values []any) (*schema.Resource, error) {
	table := c.Tables.Get(tableName)
	resource := schema.NewResourceData(table, nil, values)
	for i, col := range table.Columns {
		v, err := prepareValueForResourceSet(col, values[i])
		if err != nil {
			return nil, err
		}
		if err := resource.Set(col.Name, v); err != nil {
			return nil, err
		}
	}
	return resource, nil
}

func prepareValueForResourceSet(col schema.Column, v any) (any, error) {
	switch tp := col.Type.(type) {
	case *arrow.StringType:
		if value, ok := v.(driver.Valuer); ok {
			if value == driver.Valuer(nil) {
				v = nil
			} else {
				val, err := value.Value()
				if err != nil {
					return nil, err
				}
				if s, ok := val.(string); ok {
					v = s
				}
			}
		}
	case *arrow.Time32Type:
		t, err := v.(pgtype.Time).TimeValue()
		if err != nil {
			return nil, err
		}
		v = stringForTime(t, tp.Unit)
	case *arrow.Time64Type:
		t, err := v.(pgtype.Time).TimeValue()
		if err != nil {
			return nil, err
		}
		v = stringForTime(t, tp.Unit)
	}
	return v, nil
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
