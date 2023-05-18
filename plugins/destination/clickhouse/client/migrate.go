package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	have, err := c.getTableDefinitions(ctx, tables)
	if err != nil {
		return err
	}

	want, err := typeconv.CanonizedTables(tables)
	if err != nil {
		return err
	}
	if c.mode != specs.MigrateModeForced {
		unsafe := unsafeSchemaChanges(have, want)
		if len(unsafe) > 0 {
			return fmt.Errorf("'migrate_mode: forced' is required for the following changes: \n%s", util.SchemasChangesPrettified(unsafe))
		}
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, want := range want {
		want := want
		eg.Go(func() (err error) {
			c.logger.Info().Str("table", want.Name).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", want.Name).Msg("Migrating table done")
			}()
			if len(want.Columns) == 0 {
				c.logger.Warn().Str("table", want.Name).Msg("Table with no columns, skip")
				return nil
			}

			have := have.Get(want.Name)
			if have == nil {
				return c.createTable(ctx, want)
			}

			return c.autoMigrate(ctx, have, want)
		})
	}

	return eg.Wait()
}

func unsafeSchemaChanges(have, want schema.Tables) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, w := range want {
		current := have.Get(w.Name)
		if current == nil {
			continue
		}
		unsafe := unsafeChanges(w.GetChanges(current))
		if len(unsafe) > 0 {
			result[w.Name] = unsafe
		}
	}
	return result
}

func unsafeChanges(changes []schema.TableColumnChange) []schema.TableColumnChange {
	unsafe := make([]schema.TableColumnChange, 0, len(changes))
	for _, c := range changes {
		if needsTableDrop(c) {
			unsafe = append(unsafe, c)
		}
	}
	return slices.Clip(unsafe)
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")

	query, err := queries.CreateTable(table, c.spec.Cluster, c.spec.Engine)
	if err != nil {
		return err
	}

	if err := c.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("failed to create table, query:\n%s\nerror: %w", query, err)
	}
	return nil
}

func (c *Client) dropTable(ctx context.Context, table *schema.Table) error {
	c.logger.Debug().Str("table", table.Name).Msg("Dropping table")

	return c.conn.Exec(ctx, queries.DropTable(table, c.spec.Cluster))
}

func needsTableDrop(change schema.TableColumnChange) bool {
	// We can safely add a nullable column without dropping the table
	if change.Type == schema.TableColumnChangeTypeAdd && !change.Current.NotNull {
		return false
	}

	// We can safely ignore removal of nullable columns without dropping the table
	if change.Type == schema.TableColumnChangeTypeRemove && !change.Previous.NotNull {
		return false
	}

	// TODO: add check for update + new type is extending the current type (uint8 -> uint16, float32 -> float64, new struct field, etc).
	return true
}

func (c *Client) autoMigrate(ctx context.Context, have, want *schema.Table) error {
	changes := want.GetChanges(have)

	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		// we can get here only with migrate_mode: forced
		if err := c.dropTable(ctx, have); err != nil {
			return err
		}

		return c.createTable(ctx, want)
	}

	for _, change := range changes {
		// we only handle new columns
		if change.Type != schema.TableColumnChangeTypeAdd {
			continue
		}

		c.logger.Debug().Str("table", want.Name).Str("column", change.Current.Name).Msg("Adding new column")

		query, err := queries.AddColumn(want.Name, c.spec.Cluster, change.Current)
		if err != nil {
			return err
		}

		err = c.conn.Exec(ctx, query)
		if err != nil {
			return err
		}
	}

	return nil
}
