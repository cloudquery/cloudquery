package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) ensureTVP(ctx context.Context, table *schema.Table) error {
	if !c.pkEnabled() {
		return nil
	}

	_, err := c.db.ExecContext(ctx, queries.TVPDrop(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to drop TVP proc & type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPType(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPProc(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP proc for table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) insertTVP(ctx context.Context, table *schema.Table, data [][]any) error {
	query, params := queries.TVPQuery(c.schemaName, table, data)
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query, params...)
		return err
	})
}
