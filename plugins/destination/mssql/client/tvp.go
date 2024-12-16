package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func useTVP(table *schema.Table) bool {
	return len(table.PrimaryKeys()) > 0
}

func (c *Client) ensureTVP(ctx context.Context, table *schema.Table) (err error) {
	if !useTVP(table) {
		return nil
	}

	query, params := queries.TVPDropProc(c.spec.Schema, table)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP proc for table %s: %w", table.Name, err)
	}

	query, params = queries.TVPDropType(c.spec.Schema, table)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddType(c.spec.Schema, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddProc(c.spec.Schema, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP proc for table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) insertTVP(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	query, params, err := queries.TVPQuery(c.spec.Schema, table, records)
	if err != nil {
		return err
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query, params...)
		return err
	})
}
