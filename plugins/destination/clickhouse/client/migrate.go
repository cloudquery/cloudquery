package client

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

// MigrateTables relies on the CLI/client to lock before running migration.
func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	have, err := c.getTableDefinitions(ctx, messages)
	if err != nil {
		return err
	}

	want, err := typeconv.CanonizedTables(messages)
	if err != nil {
		return err
	}

	if err := c.checkForced(have, want, messages); err != nil {
		return err
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
				return c.createTable(ctx, want, c.spec.PartitionBySyncGroupID)
			}

			return c.autoMigrate(ctx, have, want, c.spec.PartitionBySyncGroupID)
		})
	}

	return eg.Wait()
}

func (c *Client) checkForced(have, want schema.Tables, messages message.WriteMigrateTables) error {
	forcedErr := false
	for _, m := range messages {
		if m.MigrateForce {
			continue
		}

		// check that this migration can go through
		have := have.Get(m.Table.Name)
		if have == nil {
			continue // create new is always OK
		}
		want := want.Get(m.Table.Name) // and it should never be nil
		if unsafe := unsafeChanges(want.GetChanges(have)); len(unsafe) > 0 {
			c.logger.Error().
				Str("table", m.Table.Name).
				Str("changes", util.ChangesPrettified(m.Table.Name, unsafe)).
				Msg("migrate manually or consider using 'migrate_mode: forced'")
			forcedErr = true
		}
	}

	if forcedErr {
		return errors.New("migrate manually or consider using 'migrate_mode: forced'")
	}
	return nil
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

func (c *Client) createTable(ctx context.Context, table *schema.Table, partitionBySyncGroupID bool) (err error) {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")

	query, err := queries.CreateTable(table, c.spec.Cluster, c.spec.Engine, partitionBySyncGroupID)
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

func (c *Client) autoMigrate(ctx context.Context, have, want *schema.Table, partitionBySyncGroupID bool) error {
	changes := want.GetChanges(have)

	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		// we can get here only with migrate_mode: forced
		if err := c.dropTable(ctx, have); err != nil {
			return err
		}

		return c.createTable(ctx, want, partitionBySyncGroupID)
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
