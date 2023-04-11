package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func (c *Client) schemaTables(ctx context.Context, tables schema.Tables) (schema.Tables, error) {
	query, params := queries.AllTables(c.schemaName)
	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	schemaTables := make(schema.Tables, 0)
	if err := processRows(rows, func(row *sql.Rows) error {
		var tableCatalog string
		var tableName string
		var tableType string
		var schemaType string

		if err := row.Scan(&tableCatalog, &tableType, &tableName, &schemaType); err != nil {
			return err
		}
		if tables.Get(tableName) == nil {
			return nil
		}
		schemaTables = append(schemaTables, &schema.Table{Name: tableName})
		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	for _, table := range schemaTables {
		columns, err := c.getTableColumns(ctx, table)
		if err != nil {
			return nil, err
		}
		table.Columns = columns
		pks, err := c.getTablePK(ctx, table)
		if err != nil {
			return nil, err
		}
		for _, pk := range pks {
			if col := table.Columns.Get(pk); col != nil {
				col.CreationOptions.PrimaryKey = true
			}
		}
	}

	return schemaTables, nil
}

func (c *Client) normalizedTables(tables schema.Tables) schema.Tables {
	var normalized schema.Tables
	for _, table := range tables.FlattenTables() {
		table := table
		for i := range table.Columns {
			// Since multiple schema types can map to the same MSSQL type we need to normalize them to avoid false positives when detecting schema changes
			// This should never return an error
			typ, _ := queries.SchemaType(table.Name, table.Columns[i].Name, queries.SQLType(table.Columns[i].Type))
			table.Columns[i].Type = typ
		}
		// If there are no PKs, we use CqID as PK
		pks := table.PrimaryKeys()
		if !c.pkEnabled() || len(pks) == 0 {
			table.Columns.Get(schema.CqIDColumn.Name).CreationOptions.PrimaryKey = true
		}

		normalized = append(normalized, table)
	}

	return normalized
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, schemaTables schema.Tables) (names []string, changes [][]schema.TableColumnChange) {
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		schemaTable := schemaTables.Get(t.Name)
		if schemaTable == nil {
			continue
		}
		changes := t.GetChanges(schemaTable)
		if !c.canAutoMigrate(changes) {
			names = append(names, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return names, tableChanges
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && (change.Current.CreationOptions.PrimaryKey || change.Current.CreationOptions.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeRemove && (change.Previous.CreationOptions.PrimaryKey || change.Previous.CreationOptions.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeUpdate {
			return false
		}
	}
	return true
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

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	var statements []string
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			def := queries.GetDefinition(&change.Current, c.pkEnabled())
			statements = append(statements, queries.AddColumn(c.schemaName, table, def))
		}
	}

	err := c.execStatements(ctx, table, statements)
	if err != nil {
		return err
	}
	return c.ensureTVP(ctx, table)
}

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	schemaTables, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	normalizedTables := c.normalizedTables(tables)

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigratableTables, changes := c.nonAutoMigratableTables(normalizedTables, schemaTables)
		if len(nonAutoMigratableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigratableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}
		schemaTable := schemaTables.Get(table.Name)
		if schemaTable == nil {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, table); err != nil {
				return err
			}
			continue
		}

		changes := table.GetChanges(schemaTable)
		if c.canAutoMigrate(changes) {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table, changes); err != nil {
				return err
			}
			continue
		}

		c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
		if err := c.recreateTable(ctx, table); err != nil {
			return err
		}
	}

	return nil
}
