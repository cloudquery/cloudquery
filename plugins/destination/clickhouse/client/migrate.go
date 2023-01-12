package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	flattened := tables.FlattenTables()

	current, err := c.getTableDefinitions(ctx)
	if err != nil {
		return err
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, def := range queries.GetTableDefinitions(flattened) {
		def := def
		eg.Go(func() (err error) {
			c.logger.Info().Str("table", def.Name).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", def.Name).Msg("Migrating table done")
			}()

			if len(def.Columns) == 0 {
				c.logger.Warn().Str("table", def.Name).Msg("Table with no columns, skip")
				return nil
			}

			curr := current[def.Name]
			if curr == nil {
				return c.createTable(ctx, def)
			}

			return c.autoMigrateTable(ctx, curr, def)
		})
	}

	return eg.Wait()
}

func (c *Client) createTable(ctx context.Context, def *queries.TableDefinition) (err error) {
	c.logger.Debug().Str("table", def.Name).Msg("Table doesn't exist, creating")

	return c.conn.Exec(ctx, queries.CreateTable(def))
}

func (c *Client) autoMigrateTable(ctx context.Context, current, definition *queries.TableDefinition) (err error) {
	tableName := definition.Name
	c.logger.Info().Str("table", tableName).Msg("Table exists, migrating")

	for _, column := range definition.GetAddedColumns(current) {
		if err := c.conn.Exec(ctx, queries.AddColumn(tableName, column)); err != nil {
			return err
		}
	}

	for _, column := range definition.GetChangedColumns(current) {
		if err := c.conn.Exec(ctx, queries.ModifyColumn(tableName, column)); err != nil {
			return err
		}
	}

	return nil
}
