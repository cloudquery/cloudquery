package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) useTVP(table *schema.Table) bool {
	return c.pkEnabled() && len(table.PrimaryKeys()) > 0
}

func (c *Client) ensureTVP(ctx context.Context, table *schema.Table) (err error) {
	if !c.useTVP(table) {
		return nil
	}

	query, params := queries.TVPDropProc(c.schemaName, table)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP proc for table %s: %w", table.Name, err)
	}

	query, params = queries.TVPDropType(c.schemaName, table)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddType(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddProc(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP proc for table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) insertTVP(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	query, params, err := queries.TVPQuery(c.schemaName, table, records)
	if err != nil {
		return err
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query, params...)
		return err
	})
}
