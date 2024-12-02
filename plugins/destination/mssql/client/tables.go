package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) createTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	_, err = c.db.ExecContext(ctx, queries.CreateTable(c.spec.Schema, table))
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}

	return c.ensureTVP(ctx, table)
}

func (c *Client) dropTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table exists, dropping")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	_, err = c.db.ExecContext(ctx, queries.DropTable(c.spec.Schema, table))
	if err != nil {
		return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) recreateTable(ctx context.Context, table *schema.Table) error {
	err := c.dropTable(ctx, table)
	if err != nil {
		return err
	}

	return c.createTable(ctx, table)
}
