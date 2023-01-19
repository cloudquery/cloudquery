package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/slices"
)

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")

	pkPresent, err := c.getTablePK(ctx, table)
	if err != nil {
		return err
	}

	stalePks := c.getStalePks(table, pkPresent)
	if len(stalePks) > 0 {
		dropConstraintSQL := queries.DropPK(c.schemaName, table)
		sep := strings.Repeat("-", len(dropConstraintSQL)+1)
		query := fmt.Sprintf("%s\n%s;\n%s\n%s", sep, dropConstraintSQL, c.getDropNotNullQuery(table, stalePks), sep)
		return fmt.Errorf(
			`the following primary keys were removed from the schema %q for table %q.
You can migrate the table manually by running:
%s`, stalePks, table.Name, query)
	}

	if err := c.ensureColumns(ctx, table, pkPresent); err != nil {
		return err
	}
	return c.ensureTVP(ctx, table)
}

func (c *Client) ensureColumns(ctx context.Context, table *schema.Table, pkPresent []string) (err error) {
	current, err := c.getTableColumns(ctx, table)
	if err != nil {
		return err
	}
	c.logger.Info().Interface("defs", current).Msg("current cols")

	recreatePK := false

	var updated bool
	defer func() {
		if err == nil && !updated {
			c.logger.Info().Str("table", table.Name).Msg("Table is up to date, no changes were made")
		}
	}()

	var statements []string
	pkEnabled := c.pkEnabled()
	for _, column := range table.Columns {
		def := queries.GetDefinition(&column, pkEnabled)
		switch curr := current.Get(column.Name); {
		case curr == nil:
			// column doesn't exist
			c.logger.Info().
				Str("table", table.Name).
				Str("column", column.Name).
				Msg("Column doesn't exist, creating")

			recreatePK = recreatePK || column.CreationOptions.PrimaryKey

			statements = append(statements, queries.AddColumn(c.schemaName, table, def))
			updated = true
		case curr.Type() != def.Type():
			// column exists but type is different
			c.logger.Info().
				Str("table", table.Name).
				Str("column", column.Name).
				Str("old_type", curr.Type()).
				Str("new_type", def.Type()).
				Msg("Column exists but type is different, re-creating")

			// we need t check if the existing col is part of pk
			// if the new PK contains this column we will need to recreate the primary key
			recreatePK = recreatePK || column.CreationOptions.PrimaryKey

			// right now we will drop the column and re-create. in the future we will have an option to automigrate
			statements = append(statements, queries.DropColumn(c.schemaName, table, def))
			statements = append(statements, queries.AddColumn(c.schemaName, table, def))
			updated = true
		}

		// column exists and type is the same but constraints might differ
		// check that PK contains both of new & old columns or neither of them
		if column.CreationOptions.PrimaryKey != slices.Contains(pkPresent, column.Name) {
			recreatePK = true

			if c.pkEnabled() {
				c.logger.Info().
					Str("table", table.Name).
					Str("column", column.Name).
					Bool("pk", column.CreationOptions.PrimaryKey).
					Msg("Column exists with different primary keys")
			}
			updated = true
		}
	}

	if c.pkEnabled() && recreatePK {
		statements = append([]string{queries.DropPK(c.schemaName, table)}, statements...)
		statements = append(statements, queries.AddPK(c.schemaName, table))
	}

	return c.execStatements(ctx, table, statements)
}

func (c *Client) execStatements(ctx context.Context, table *schema.Table, statements []string) error {
	if len(statements) == 0 {
		return nil
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		query := strings.Join(statements, "\n")
		c.logger.Debug().Str("table", table.Name).Str("query", query).Msg("exec migration statement")
		_, err := c.db.ExecContext(ctx, query)
		return err
	})
}
