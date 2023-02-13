package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	current, err := c.getTableDefinitions(ctx)
	if err != nil {
		return err
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, table := range queries.NormalizedTables(tables) {
		table := table
		eg.Go(func() (err error) {
			c.logger.Info().Str("table", table.Name).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", table.Name).Msg("Migrating table done")
			}()

			if len(table.Columns) == 0 {
				c.logger.Warn().Str("table", table.Name).Msg("Table with no columns, skip")
				return nil
			}

			curr := current.Get(table.Name)
			if curr == nil {
				return c.createTable(ctx, table)
			}

			return c.autoMigrateTable(ctx, curr, table)
		})
	}

	return eg.Wait()
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")

	return c.conn.Exec(ctx, queries.CreateTable(table))
}

func (c *Client) autoMigrateTable(ctx context.Context, current, table *schema.Table) (err error) {
	tableName := table.Name
	c.logger.Info().Str("table", tableName).Msg("Table exists, migrating")

	have, want := current.GetChangedColumns(table)
	if len(have) > 0 && c.spec.MigrateMode != specs.MigrateModeForced {
		return fmt.Errorf("table %s has different types for columns %v but schema wants %v , use --force to drop the columns", table.Name, have, want)
	}

	for _, column := range want {
		if err := c.conn.Exec(ctx, queries.DropColumn(tableName, &column)); err != nil {
			return err
		}
		if err := c.conn.Exec(ctx, queries.AddColumn(tableName, &column)); err != nil {
			return err
		}
	}

	for _, column := range table.GetAddedColumns(current) {
		if err := c.conn.Exec(ctx, queries.AddColumn(tableName, &column)); err != nil {
			return err
		}
	}

	return nil
}
