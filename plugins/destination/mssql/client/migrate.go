package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	have, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	want := c.normalizedTables(tables)

	if c.spec.MigrateMode != specs.MigrateModeForced {
		unsafe := unsafeSchemaChanges(have, want)
		if len(unsafe) > 0 {
			return fmt.Errorf("the following changes cannot be automatically migrated:\n%s\n\nPlease perform this migration manually or enable `migrate_mode: 'forced'`", prettifyChanges(unsafe))
		}
	}

	for _, want := range want {
		c.logger.Info().Str("table", want.Name).Msg("Migrating table")
		if len(want.Columns) == 0 {
			c.logger.Info().Str("table", want.Name).Msg("Table with no columns, skipping")
			continue
		}

		have := have.Get(want.Name)
		if have == nil {
			c.logger.Debug().Str("table", want.Name).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, want); err != nil {
				return err
			}
			continue
		}

		c.logger.Info().Str("table", want.Name).Msg("Table exists, auto-migrating")
		if err := c.autoMigrateTable(ctx, have, want); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) autoMigrateTable(ctx context.Context, have, want *schema.Table) error {
	changes := want.GetChanges(have)
	if len(changes) == 0 {
		c.logger.Info().Str("table", want.Name).Msg("Table schema is up-to-date, skip")
		return nil
	}

	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		// we can get here only with migrate_mode: forced
		c.logger.Info().Str("table", want.Name).Msg("Table exists, force migration required")
		return c.recreateTable(ctx, want)
	}

	statements := make([]string, 0, len(changes))
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			statements = append(statements, queries.AddColumn(c.schemaName, want, &change.Current))
		}
	}

	err := c.execStatements(ctx, want.Name, statements)
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
