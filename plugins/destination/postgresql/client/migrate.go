package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	// this returns the following table in sorted manner:
	// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
	// | ordinal_position | table_name | column_name | data_type | is_primary_key| not_null  | pk_constraint_name  |
	// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
	// |              1 | users       | id          | bigint     | YES           | true 		 | cq_users_pk 	  	   |
	// |              2 | users       | name        | text       | NO            | false 	   | 					           |
	// |              3 | users       | email       | text       | NO            | true 		 | cq_users_pk         |
	// |              1 | posts       | id          | bigint     | YES           | true 		 | cq_posts_pk			   |
	// |              2 | posts       | title       | text       | NO            | false 	   | 					           |
	selectAllTables = `
SELECT
	columns.ordinal_position AS ordinal_position,
	pg_class.relname AS table_name,
	pg_attribute.attname AS column_name,
	pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod) AS data_type,
	CASE 
		WHEN conkey IS NOT NULL AND array_position(conkey, pg_attribute.attnum) > 0 THEN true
		ELSE false
	END AS is_primary_key,
	CASE 
		WHEN pg_attribute.attnotnull THEN true
		ELSE false
	END AS not_null,
	COALESCE(pg_constraint.conname, '') AS primary_key_constraint_name
FROM
	pg_catalog.pg_attribute
	INNER JOIN
	pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
	INNER JOIN
	pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
	LEFT JOIN
	pg_catalog.pg_constraint ON pg_constraint.conrelid = pg_attribute.attrelid
	AND conkey IS NOT NULL AND array_position(conkey, pg_attribute.attnum) > 0
	AND contype = 'p'
	INNER JOIN
	information_schema.columns ON columns.table_name = pg_class.relname AND columns.column_name = pg_attribute.attname AND columns.table_schema = pg_catalog.pg_namespace.nspname
WHERE
	pg_attribute.attnum > 0
	AND NOT pg_attribute.attisdropped
	AND pg_catalog.pg_namespace.nspname = '%s'
ORDER BY
	table_name ASC , ordinal_position ASC;
`
)

func (c *Client) listPgTables(ctx context.Context, pluginTables schema.Tables) (schema.Tables, error) {
	var tables schema.Tables
	rows, err := c.conn.Query(ctx, fmt.Sprintf(selectAllTables, c.currentSchemaName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ordinalPosition int
		var tableName, columnName, columnType, pkName string
		var isPrimaryKey, notNull bool
		if err := rows.Scan(&ordinalPosition, &tableName, &columnName, &columnType, &isPrimaryKey, &notNull, &pkName); err != nil {
			return nil, err
		}
		// We don't want to migrate tables that are not a part of the spec, or non CloudQuery tables
		if pluginTables.Get(tableName) == nil {
			continue
		}
		if ordinalPosition == 1 {
			tables = append(tables, &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			})
		}
		table := tables[len(tables)-1]
		if pkName != "" {
			table.PkConstraintName = pkName
		}
		table.Columns = append(table.Columns, schema.Column{
			Name: columnName,
			CreationOptions: schema.ColumnCreationOptions{
				PrimaryKey: isPrimaryKey,
				NotNull:    notNull,
			},
			Type: c.PgToSchemaType(columnType),
		})
	}
	return tables, nil
}

func (*Client) normalizeTableCockroach(table *schema.Table) *schema.Table {
	for i := range table.Columns {
		switch table.Columns[i].Type {
		case schema.TypeCIDR:
			table.Columns[i].Type = schema.TypeInet
		case schema.TypeCIDRArray:
			table.Columns[i].Type = schema.TypeInetArray
		case schema.TypeMacAddr:
			table.Columns[i].Type = schema.TypeString
		case schema.TypeMacAddrArray:
			table.Columns[i].Type = schema.TypeStringArray
		}
	}
	return table
}

func (c *Client) normalizeTable(table *schema.Table) *schema.Table {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.normalizeTableCockroach(table)
	case pgTypePostgreSQL:
		return table
	default:
		panic("unknown pg type")
	}
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	for _, table := range tables.FlattenTables() {
		table := c.normalizeTable(table)
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}

		// In postgres if column is primary key, it can't be null
		if c.enabledPks() {
			for i := range table.Columns {
				if table.Columns[i].CreationOptions.PrimaryKey {
					table.Columns[i].CreationOptions.NotNull = true
				}
			}
			// this is for backward compatibility as I believe this is as of right now defined on the source
			if len(table.PrimaryKeys()) == 0 {
				cqIdColumn := table.Columns.Get(schema.CqIDColumn.Name)
				if cqIdColumn != nil {
					cqIdColumn.CreationOptions.PrimaryKey = true
				}
			}
		}

		pgTable := pgTables.Get(table.Name)
		if pgTable != nil {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, pgTable, table); err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
	}
	conn, err := c.conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()
	if err := conn.Conn().DeallocateAll(ctx); err != nil {
		return fmt.Errorf("failed to deallocate all prepared statements: %w", err)
	}
	return nil
}

func (c *Client) alterColumn(ctx context.Context, tableName string, column schema.Column) error {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.alterColumnCockroachDB(ctx, tableName, column)
	case pgTypePostgreSQL:
		return c.alterColumnPg(ctx, tableName, column)
	default:
		return fmt.Errorf("unknown pg type: %v", c.pgType)
	}
}

func (c *Client) alterColumnCockroachDB(ctx context.Context, tableName string, column schema.Column) error {
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	// try alter column in place first
	sql := "alter table " + tableName + " alter column " + columnName + " type " + columnType
	if column.CreationOptions.NotNull {
		sql += ", ALTER " + columnName + " set not null"
	}
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		c.logger.Warn().Err(err).Str("table", tableName).Str("column", column.Name).Msg("Column type changed in place failed.")
	} else {
		return nil
	}

	sql = "alter table " + tableName + " drop column " + columnName
	// right now we will drop the column and re-create. in the future we will have an option to automigrate
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop column %s.%s: %w", column.Name, tableName, err)
	}
	sql = "alter table " + tableName + " add column " + columnName + " " + columnType
	if column.CreationOptions.NotNull {
		sql += " not null"
	}
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			return fmt.Errorf("failed to add column %s.%s: %w", tableName, column.Name, err)
		}
		// this is a weird CockroachDB error code that means we need to retry.
		if pgErr.Code != "55000" {
			return fmt.Errorf("failed to add column %s.%s with pgerror %s: %w", tableName, column.Name, pgErrToStr(pgErr), err)
		}
		if _, err := c.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to retry adding column %s.%s with pgerror %s: %w", tableName, column.Name, pgErrToStr(pgErr), err)
		}
	}
	return nil
}

func (c *Client) alterColumnPg(ctx context.Context, tableName string, column schema.Column) error {
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	// try alter column in place first
	sql := "alter table " + tableName + " alter column " + columnName + " type " + columnType
	if column.CreationOptions.NotNull {
		sql += ", ALTER " + columnName + " set not null"
	}
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		c.logger.Warn().Err(err).Str("table", tableName).Str("column", column.Name).Msg("Column type changed in place failed.")
	} else {
		return nil
	}
	tx, err := c.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to start alter column transaction: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if !errors.Is(err, pgx.ErrTxClosed) {
				c.logger.Error().Err(err).Msg("Failed to rollback alter column transaction")
			}
		}
	}()

	sql = "alter table " + tableName + " drop column " + columnName
	// right now we will drop the column and re-create. in the future we will have an option to automigrate
	if _, err := tx.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop column %s.%s: %w", column.Name, tableName, err)
	}
	sql = "alter table " + tableName + " add column " + columnName + " " + columnType
	if column.CreationOptions.NotNull {
		sql += " not null"
	}
	if _, err := tx.Exec(ctx, sql); err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			return fmt.Errorf("failed to add column %s.%s: %w", tableName, column.Name, err)
		}
		return fmt.Errorf("failed to add column %s.%s with pgerror %s: %w", tableName, column.Name, pgErrToStr(pgErr), err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit alter column transaction: %w", err)
	}
	return nil
}

func (c *Client) addColumn(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Column doesn't exist, creating")
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	sql := "alter table " + tableName + " add column " + columnName + " " + columnType
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) alterPKConstraint(ctx context.Context, pgTable *schema.Table, table *schema.Table) error {
	c.logger.Info().Str("table", table.Name).Msg("Recreating primary keys")
	tableName := pgx.Identifier{table.Name}.Sanitize()

	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction to recreate primary keys: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if !errors.Is(err, pgx.ErrTxClosed) {
				c.logger.Error().Err(err).Msg("failed to rollback transaction")
			}
		}
	}()
	dropConstraintSQL := "alter table " + tableName + " drop constraint if exists " + pgTable.PkConstraintName
	if _, err := tx.Exec(ctx, dropConstraintSQL); err != nil {
		return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
	}

	sql := "alter table " + tableName + " add primary key (" + strings.Join(table.PrimaryKeys(), ",") + ")"
	if _, err := tx.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction to recreate primary keys: %w", err)
	}
	return nil
}

func (c *Client) autoMigrateTable(ctx context.Context, pgTable *schema.Table, table *schema.Table) error {
	changedColumns, changedPgColumns := table.GetChangedColumns(pgTable)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		if c.enabledPks() && !table.IsPrimaryKeyEqual(pgTable) {
			return fmt.Errorf("postgres table %s has different primary keys %v but schema wants %v , add `migrate_mode: forced` to the destination spec to drop the column", table.Name, pgTable.PrimaryKeys(), table.PrimaryKeys())
		}
		if changedColumns != nil {
			return fmt.Errorf("postgres table %s has different types for columns %v but schema wants %v , add `migrate_mode: forced` to the destination spec to drop the column", table.Name, changedPgColumns, changedColumns)
		}
	}

	for _, col := range changedColumns {
		if err := c.alterColumn(ctx, table.Name, col); err != nil {
			return err
		}
	}

	if addedColumns := table.GetAddedColumns(pgTable); addedColumns != nil {
		for _, col := range addedColumns {
			if err := c.addColumn(ctx, table.Name, col); err != nil {
				return err
			}
		}
	}

	if c.enabledPks() && !table.IsPrimaryKeyEqual(pgTable) {
		if err := c.alterPKConstraint(ctx, pgTable, table); err != nil {
			return err
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
		if col.CreationOptions.NotNull {
			fieldDef += " NOT NULL"
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
	}
	sb.WriteString(")")
	_, err := c.conn.Exec(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}
	return nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
