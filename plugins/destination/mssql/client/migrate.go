package client

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// MigrateTables relies on the CLI/client to lock before running migration.
func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	have, err := c.schemaTables(ctx, messages)
	if err != nil {
		return err
	}

	want := normalizedTables(messages)

	if err := c.checkForced(have, want, messages); err != nil {
		return err
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
				Str("changes", prettifyChanges(m.Table.Name, unsafe)).
				Msg("migrate manually or consider using 'migrate_mode: forced'")
			forcedErr = true
		}
	}

	if forcedErr {
		return errors.New("migrate manually or consider using 'migrate_mode: forced'")
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
			statements = append(statements, queries.AddColumn(c.spec.Schema, want, &change.Current))
		}
		if change.Type == schema.TableColumnChangeTypeUpdate {
			statements = append(statements, queries.UpdateColumnType(c.spec.Schema, want, &change.Current))
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
