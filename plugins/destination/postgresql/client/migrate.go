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

	selectAllTablesCockroach = `
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
	AND information_schema.columns.is_hidden != 'YES'
ORDER BY
	table_name ASC , ordinal_position ASC;
`
)

func (c *Client) listPgTables(ctx context.Context, pluginTables schema.Tables) (schema.Tables, error) {
	var tables schema.Tables
	sql := selectAllTables
	if c.pgType == pgTypeCockroachDB {
		sql = selectAllTablesCockroach
	}
	rows, err := c.conn.Query(ctx, fmt.Sprintf(sql, c.currentSchemaName))
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
		schemaType, err := c.PgToSchemaType(tableName, columnName, columnType)
		if err != nil {
			return nil, err
		}
		table.Columns = append(table.Columns, schema.Column{
			Name: columnName,
			CreationOptions: schema.ColumnCreationOptions{
				PrimaryKey: isPrimaryKey,
				NotNull:    notNull,
			},
			Type: schemaType,
		})
	}
	return tables, nil
}

func (c *Client) normalizeTableCockroach(table *schema.Table) *schema.Table {
	for i := range table.Columns {
		if !c.enabledPks() {
			table.Columns[i].CreationOptions.PrimaryKey = false
		}
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
		if table.Columns[i].CreationOptions.PrimaryKey {
			table.Columns[i].CreationOptions.NotNull = true
		}
	}
	return table
}

func (c *Client) normalizeTablePg(table *schema.Table) *schema.Table {
	for i := range table.Columns {
		if !c.enabledPks() {
			table.Columns[i].CreationOptions.PrimaryKey = false
		}
		if table.Columns[i].CreationOptions.PrimaryKey {
			table.Columns[i].CreationOptions.NotNull = true
		}
		// this is for backward compatibility as I believe this is as of right now defined on the source
		if c.enabledPks() && len(table.PrimaryKeys()) == 0 {
			cqIdColumn := table.Columns.Get(schema.CqIDColumn.Name)
			if cqIdColumn != nil {
				cqIdColumn.CreationOptions.PrimaryKey = true
			}
		}
	}
	return table
}

func (c *Client) normalizeTable(table *schema.Table) *schema.Table {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.normalizeTableCockroach(table)
	case pgTypePostgreSQL:
		return c.normalizeTablePg(table)
	default:
		panic("unknown pg type")
	}
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if err := c.addColumn(ctx, table.Name, change.Current); err != nil {
				return err
			}
		case schema.TableColumnChangeTypeRemove:
			continue
		default:
			panic("unknown change type")
		}
	}
	return nil
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if change.Current.CreationOptions.PrimaryKey || change.Current.CreationOptions.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.CreationOptions.PrimaryKey || change.Previous.CreationOptions.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeUpdate:
			return false
		default:
			panic("unknown change type")
		}
	}
	return true
}

// normalize the requested schema to be compatible with what Postgres supports
func (c *Client) normalizeTables(tables schema.Tables) schema.Tables {
	flattenedTables := tables.FlattenTables()
	for _, table := range flattenedTables {
		c.normalizeTable(table)
	}
	return flattenedTables
}

func (c *Client) nonAutoMigrableTables(tables schema.Tables, pgTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		pgTable := pgTables.Get(t.Name)
		if pgTable == nil {
			continue
		}
		changes := t.GetChanges(pgTable)
		if !c.canAutoMigrate(changes) {
			result = append(result, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	tables = c.normalizeTables(tables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(tables, pgTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range tables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}
		pgTable := pgTables.Get(table.Name)
		if pgTable == nil {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		} else {
			changes := table.GetChanges(pgTable)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, table.Name); err != nil {
					return err
				}
				if err := c.createTableIfNotExist(ctx, table); err != nil {
					return err
				}
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

func (c *Client) dropTable(ctx context.Context, tableName string) error {
	c.logger.Info().Str("table", tableName).Msg("Dropping table")
	sql := "drop table " + tableName
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
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
		if col.CreationOptions.Unique {
			fieldDef += " UNIQUE"
		}
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
