package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, scs schema.Schemas) error {
	have, err := c.schemaTables(ctx, scs)
	if err != nil {
		return err
	}

	want := c.normalizedSchemas(scs)

	if c.spec.MigrateMode != specs.MigrateModeForced {
		unsafe := unsafeSchemaChanges(have, want)
		if len(unsafe) > 0 {
			return fmt.Errorf("'migrate_mode: forced' is required for the following changes: \n%s", prettifyChanges(unsafe))
		}
	}

	for _, want := range want {
		tableName := schema.TableName(want)
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(want.Fields()) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}

		have := have.SchemaByName(tableName)
		if have == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, want); err != nil {
				return err
			}
			continue
		}

		c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
		if err := c.autoMigrateTable(ctx, have, want); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) autoMigrateTable(ctx context.Context, have, want *arrow.Schema) error {
	tableName := schema.TableName(want)
	changes := schema.GetSchemaChanges(want, have)

	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		// we can get here only with migrate_mode: forced
		c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
		return c.recreateTable(ctx, want)
	}

	statements := make([]string, 0, len(changes))
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			def := queries.GetDefinition(change.Current, c.pkEnabled())
			statements = append(statements, queries.AddColumn(c.schemaName, want, def))
		}
	}

	err := c.execStatements(ctx, tableName, statements)
	if err != nil {
		return err
	}

	return c.ensureTVP(ctx, want)
}

func (c *Client) execStatements(ctx context.Context, tableName string, statements []string) error {
	if len(statements) == 0 {
		return nil
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		query := strings.Join(statements, "\n")
		c.logger.Debug().Str("table", tableName).Str("query", query).Msg("exec migration statement")
		_, err := c.db.ExecContext(ctx, query)
		return err
	})
}
