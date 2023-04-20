package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"golang.org/x/sync/errgroup"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, scs schema.Schemas) error {
	currentSchema, err := c.getTableDefinitions(ctx)
	if err != nil {
		return err
	}

	newSchema, err := typeconv.CanonizedSchemas(scs)
	if err != nil {
		return err
	}
	if c.mode != specs.MigrateModeForced {
		nonSafeMigratableTables, changes := c.nonAutoMigratableTables(newSchema, currentSchema)
		if len(nonSafeMigratableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonSafeMigratableTables, ","), changes)
		}
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, sc := range newSchema {
		sc := sc
		eg.Go(func() (err error) {
			tableName := schema.TableName(sc)
			c.logger.Info().Str("table", tableName).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", tableName).Msg("Migrating table done")
			}()
			if len(sc.Fields()) == 0 {
				c.logger.Warn().Str("table", tableName).Msg("Table with no columns, skip")
				return nil
			}

			current := currentSchema.SchemaByName(tableName)
			if current == nil {
				return c.createTable(ctx, sc)
			}

			return c.autoMigrate(ctx, sc, current)
		})
	}

	return eg.Wait()
}

func (c *Client) nonAutoMigratableTables(want, have schema.Schemas) ([]string, [][]schema.FieldChange) {
	var result []string
	var tableChanges [][]schema.FieldChange
	for _, w := range want {
		current := have.SchemaByName(schema.TableName(w))
		if current == nil {
			continue
		}
		changes := schema.GetSchemaChanges(w, current)
		if canSafelyMigrate(changes) {
			result = append(result, schema.TableName(w))
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func canSafelyMigrate(changes []schema.FieldChange) bool {
	for _, change := range changes {
		needsDrop := needsTableDrop(change)
		if needsDrop {
			return false
		}
	}
	return true
}

func (c *Client) createTable(ctx context.Context, sc *arrow.Schema) (err error) {
	c.logger.Debug().Str("table", schema.TableName(sc)).Msg("Table doesn't exist, creating")

	query, err := queries.CreateTable(sc, c.spec.Cluster, c.spec.Engine)
	if err != nil {
		return err
	}

	return c.conn.Exec(ctx, query)
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

	return true
}

func (c *Client) autoMigrate(ctx context.Context, have, want *arrow.Schema) (err error) {
	changes := schema.GetSchemaChanges(want, have)

	if !canSafelyMigrate(changes) {
		// We already checked that the mode is forced, so just recreate
		err := c.dropTable(ctx, have)
		if err != nil {
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
