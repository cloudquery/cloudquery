package client

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	currentSchema, err := c.getTableDefinitions(ctx)
	if err != nil {
		return err
	}

	newSchema := queries.NormalizedTables(tables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonSafeMigratableTables, changes := c.nonAutoMigrableTables(newSchema, currentSchema)
		if len(nonSafeMigratableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonSafeMigratableTables, ","), changes)
		}
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, table := range newSchema {
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

			current := currentSchema.Get(table.Name)
			if current == nil {
				return c.createTable(ctx, table)
			}

			return c.autoMigrate(ctx, table, current)
		})
	}

	return eg.Wait()
}

func (c *Client) nonAutoMigrableTables(tables schema.Tables, currentTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		current := currentTables.Get(t.Name)
		if current == nil {
			continue
		}
		changes := t.GetChanges(current)
		if !c.canSafelyMigrate(changes) {
			result = append(result, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func (*Client) canSafelyMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		needsDrop := needsTableDrop(change)
		if needsDrop {
			return false
		}
	}
	return true
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")

	return c.conn.Exec(ctx, queries.CreateTable(table))
}

func (c *Client) dropTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Dropping table")

	return c.conn.Exec(ctx, queries.DropTable(table))
}

func needsTableDrop(change schema.TableColumnChange) bool {
	// We can safely add a nullable column without dropping the table
	if change.Type == schema.TableColumnChangeTypeAdd && !change.Current.CreationOptions.NotNull {
		return false
	}

	// We can safely ignore removal of nullable columns without dropping the table
	if change.Type == schema.TableColumnChangeTypeRemove && !change.Previous.CreationOptions.NotNull {
		return false
	}

	return true
}

func changesSorter(a, b schema.TableColumnChange) bool {
	return needsTableDrop(a) && !needsTableDrop(b)
}

func (c *Client) autoMigrate(ctx context.Context, table *schema.Table, current *schema.Table) (err error) {
	changes := table.GetChanges(current)
	sort.SliceStable(changes, func(i, j int) bool {
		return changesSorter(changes[i], changes[j])
	})
	for _, change := range changes {
		switch {
		case change.Type == schema.TableColumnChangeTypeAdd && !change.Current.CreationOptions.NotNull:
			c.logger.Debug().Str("table", table.Name).Str("column", change.Current.Name).Msg("Adding new column")
			err := c.conn.Exec(ctx, queries.AddColumn(table.Name, &change.Current))
			if err != nil {
				return err
			}
		case change.Type == schema.TableColumnChangeTypeRemove && !change.Previous.CreationOptions.NotNull:
			continue
		default:
			err := c.dropTable(ctx, table)
			if err != nil {
				return err
			}
			return c.createTable(ctx, table)
		}
	}

	return nil
}
