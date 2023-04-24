package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, scs schema.Schemas) error {
	have, err := c.getTableDefinitions(ctx)
	if err != nil {
		return err
	}

	want, err := typeconv.CanonizedSchemas(scs)
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
			tableName := schema.TableName(want)
			c.logger.Info().Str("table", tableName).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", tableName).Msg("Migrating table done")
			}()
			if len(want.Fields()) == 0 {
				c.logger.Warn().Str("table", tableName).Msg("Table with no columns, skip")
				return nil
			}

			have := have.SchemaByName(tableName)
			if have == nil {
				return c.createTable(ctx, want)
			}

			return c.autoMigrate(ctx, have, want)
		})
	}

	return eg.Wait()
}

func unsafeSchemaChanges(have, want schema.Schemas) map[string]schema.FieldChanges {
	result := make(map[string]schema.FieldChanges)
	for _, w := range want {
		current := have.SchemaByName(schema.TableName(w))
		if current == nil {
			continue
		}
		unsafe := unsafeChanges(schema.GetSchemaChanges(w, current))
		if len(unsafe) > 0 {
			result[schema.TableName(w)] = unsafe
		}
	}
	return result
}

func unsafeChanges(changes []schema.FieldChange) schema.FieldChanges {
	unsafe := make([]schema.FieldChange, 0, len(changes))
	for _, c := range changes {
		if needsTableDrop(c) {
			unsafe = append(unsafe, c)
		}
	}
	return slices.Clip(unsafe)
}

func (c *Client) createTable(ctx context.Context, sc *arrow.Schema) (err error) {
	c.logger.Debug().Str("table", schema.TableName(sc)).Msg("Table doesn't exist, creating")

	query, err := queries.CreateTable(sc, c.spec.Cluster, c.spec.Engine)
	if err != nil {
		return err
	}

	if err := c.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("failed to create table, query:\n%s\nerror: %w", query, err)
	}
	return nil
}

func (c *Client) dropTable(ctx context.Context, sc *arrow.Schema) error {
	c.logger.Debug().Str("table", schema.TableName(sc)).Msg("Dropping table")

	return c.conn.Exec(ctx, queries.DropTable(sc, c.spec.Cluster))
}

func needsTableDrop(change schema.FieldChange) bool {
	// We can safely add a nullable column without dropping the table
	if change.Type == schema.TableColumnChangeTypeAdd && change.Current.Nullable {
		return false
	}

	// We can safely ignore removal of nullable columns without dropping the table
	if change.Type == schema.TableColumnChangeTypeRemove && !change.Previous.Nullable {
		return false
	}

	// TODO: add check for update + new type is extending the current type (uint8 -> uint16, float32 -> float64, new struct field, etc).
	return true
}

func (c *Client) autoMigrate(ctx context.Context, have, want *arrow.Schema) error {
	changes := schema.GetSchemaChanges(want, have)

	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		// we can get here only with migrate_mode: forced
		if err := c.dropTable(ctx, have); err != nil {
			return err
		}

		return c.createTable(ctx, want)
	}

	tableName := schema.TableName(want)
	for _, change := range changes {
		// we only handle new columns
		if change.Type != schema.TableColumnChangeTypeAdd {
			continue
		}

		c.logger.Debug().Str("table", tableName).Str("column", change.Current.Name).Msg("Adding new column")

		query, err := queries.AddColumn(tableName, c.spec.Cluster, change.Current)
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
