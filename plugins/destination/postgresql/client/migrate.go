package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v5"
)

const (
	isTableExistSQL = "select count(*) from information_schema.tables where table_name = $1 and table_schema in (SELECT TRIM(BOTH from unnest(string_to_array(reset_val,','))) FROM pg_settings WHERE name='search_path')"

	// https://wiki.postgresql.org/wiki/Retrieve_primary_key_columns
	sqlSelectPrimaryKeys = `
SELECT a.attname as pkey FROM pg_index i       
JOIN   pg_attribute a ON a.attrelid = i.indrelid
  AND a.attnum = ANY(i.indkey)
WHERE  i.indrelid = $1::regclass
AND    i.indisprimary;
`
)

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}
		tableExist, err := c.isTableExistSQL(ctx, table.Name)
		if err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExist {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table); err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
		if err := c.Migrate(ctx, table.Relations); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) isTableExistSQL(ctx context.Context, table string) (bool, error) {
	var tableExist int
	if err := c.conn.QueryRow(ctx, isTableExistSQL, table).Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	return tableExist == 1, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	var err error
	var pgColumns *pgTableColumns
	var pgPKs map[string]bool

	// create the new column as it doesn't exist
	tableName := pgx.Identifier{table.Name}.Sanitize()
	if pgColumns, err = c.getPgTableColumns(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	if pgPKs, err = c.getPgTablePrimaryKeys(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s primary key columns: %w", table.Name, err)
	}

	reCreatePrimaryKeys := false

	for _, col := range table.Columns {
		columnName := pgx.Identifier{col.Name}.Sanitize()
		columnType := c.SchemaTypeToPg(col.Type)
		if columnType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type not supported, skipping")
			continue
		}
		pgColumn := pgColumns.getPgColumn(col.Name)

		switch {
		case pgColumn == nil:
			c.logger.Info().Str("table", table.Name).Str("column", col.Name).Msg("Column doesn't exist, creating")

			sql := "alter table " + tableName + " add column " + columnName + " " + columnType
			if col.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			if _, err := c.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, table.Name, err)
			}
		case pgColumn.typ != columnType:
			c.logger.Info().Str("table", table.Name).Str("column", col.Name).Str("old_type", pgColumn.typ).Str("new_type", columnType).Msg("Column exists but type is different, re-creating")
			// column exists but type is different

			// if this column contains primary key we will need to recreate the primary key
			if c.enabledPks() && col.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			sql := "alter table " + tableName + " drop column " + columnName
			// right now we will drop the column and re-create. in the future we will have an option to automigrate
			if _, err := c.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to drop column %s on table %s: %w", col.Name, table.Name, err)
			}
			sql = "alter table " + tableName + " add column " + columnName + " " + columnType
			if _, err := c.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, table.Name, err)
			}
		}

		// column exists and type is the same but constraints might differ
		if c.enabledPks() && pgPKs[col.Name] != col.CreationOptions.PrimaryKey {
			c.logger.Info().Str("table", table.Name).Str("column", col.Name).Bool("pk", col.CreationOptions.PrimaryKey).Msg("Column exists with different primary keys")
			reCreatePrimaryKeys = true
		}
	}
	if reCreatePrimaryKeys {
		c.logger.Info().Str("table", table.Name).Msg("Recreating primary keys")
		if err := c.setNullOnPks(ctx, table); err != nil {
			return fmt.Errorf("failed to enforce not null on primary keys: %w", err)
		}

		tx, err := c.conn.Begin(ctx)
		if err != nil {
			return fmt.Errorf("failed to begin transaction to recreate primary keys: %w", err)
		}
		constraintName := pgx.Identifier{table.Name + "_cqpk"}.Sanitize()
		sql := "alter table " + tableName + " drop constraint if exists " + constraintName
		if _, err := tx.Exec(ctx, sql); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				c.logger.Error().Err(err).Msg("failed to rollback transaction")
			}
			return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
		}

		sql = "alter table " + tableName + " add constraint " + constraintName + " primary key (" + strings.Join(table.PrimaryKeys(), ",") + ")"
		if _, err := tx.Exec(ctx, sql); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				c.logger.Error().Err(err).Msg("failed to rollback transaction")
			}
			return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit transaction to recreate primary keys: %w", err)
		}
	}
	return nil
}

func (c *Client) setNullOnPks(ctx context.Context, table *schema.Table) error {
	for _, col := range table.PrimaryKeys() {
		sql := "alter table " + pgx.Identifier{table.Name}.Sanitize() + " alter column " + pgx.Identifier{col}.Sanitize() + " set not null"
		if _, err := c.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to set not null on column %s on table %s: %w", col, table.Name, err)
		}
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	tableName := pgx.Identifier{table.Name}.Sanitize()
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(tableName)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		pgType := c.SchemaTypeToPg(col.Type)
		if pgType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		columnName := pgx.Identifier{col.Name}.Sanitize()
		fieldDef := columnName + " " + pgType
		if col.Name == "_cq_id" {
			// _cq_id column should always have a "unique not null" constraint
			fieldDef += " UNIQUE NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.enabledPks() && col.CreationOptions.PrimaryKey {
			primaryKeys = append(primaryKeys, col.Name)
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	} else {
		// if no primary keys are defined, add a PK constraint for _cq_id
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (_cq_id)")
	}
	sb.WriteString(")")
	_, err := c.conn.Exec(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}
	return nil
}

func (c *Client) getPgTablePrimaryKeys(ctx context.Context, tableName string) (map[string]bool, error) {
	pks := map[string]bool{}
	rows, err := c.conn.Query(ctx, sqlSelectPrimaryKeys, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		pks[strings.ToLower(column)] = true
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pks, nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
