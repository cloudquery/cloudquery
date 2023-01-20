package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) createTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	_, err = c.db.ExecContext(ctx, queries.CreateTable(c.schemaName, table, c.pkEnabled()))
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}

	return c.ensureTVP(ctx, table)
}
