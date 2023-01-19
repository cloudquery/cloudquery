package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(100)
	for _, table := range tables {
		table := table
		eg.Go(func() error {
			return c.migrate(ctx, table)
		})
	}
	return eg.Wait()
}

func (c *Client) migrate(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Info().Str("table", table.Name).Msg("Migrating table started")
	defer c.logger.Info().Str("table", table.Name).Msg("Migrating table done")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	if len(table.Columns) == 0 {
		c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
		return nil
	}

	switch tableExist, err := c.tableExists(ctx, table); {
	case err != nil:
		return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
	case tableExist:
		if err := c.autoMigrateTable(ctx, table); err != nil {
			return err
		}
	default:
		if err := c.createTable(ctx, table); err != nil {
			return err
		}
	}

	return c.Migrate(ctx, table.Relations)
}
