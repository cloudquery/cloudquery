package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
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

func (c *Client) listPgTables(ctx context.Context, pluginTables schema.Schemas) (schema.Schemas, error) {
	var tables schema.Schemas
	var fields []arrow.Field
	tableMetaData := make(map[string]string)
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
		if pluginTables.SchemaByName(tableName) == nil {
			continue
		}
		if ordinalPosition == 1 {
			if fields != nil {
				md := arrow.MetadataFrom(tableMetaData)
				tables = append(tables, arrow.NewSchema(fields, &md))
				fields = nil
				tableMetaData = make(map[string]string, 0)
			}
			tableMetaData[schema.MetadataTableName] = tableName
		}
		if pkName != "" {
			tableMetaData[schema.MetadataConstraintName] = pkName
		}
		schemaType := c.PgToSchemaType(columnType)
		fields = append(fields, arrow.Field{
			Name:     columnName,
			Type:     schemaType,
			Nullable: !notNull,
			Metadata: arrow.MetadataFrom(map[string]string{
				schema.MetadataPrimaryKey: strconv.FormatBool(isPrimaryKey),
			}),
		})
	}
	if fields != nil {
		md := arrow.MetadataFrom(tableMetaData)
		tables = append(tables, arrow.NewSchema(fields, &md))
	}
	return tables, nil
}

// func (c *Client) normalizeTableCockroach(table *arrow.Schema) *arrow.Schema {
// 	for i := range table.Fields() {
// 		if !c.enabledPks() {
// 			table.Columns[i].CreationOptions.PrimaryKey = false
// 		}
// 		switch table.Columns[i].Type {
// 		case schema.TypeCIDR:
// 			table.Columns[i].Type = schema.TypeInet
// 		case schema.TypeCIDRArray:
// 			table.Columns[i].Type = schema.TypeInetArray
// 		case schema.TypeMacAddr:
// 			table.Columns[i].Type = schema.TypeString
// 		case schema.TypeMacAddrArray:
// 			table.Columns[i].Type = schema.TypeStringArray
// 		}
// 		if table.Columns[i].CreationOptions.PrimaryKey {
// 			table.Columns[i].CreationOptions.NotNull = true
// 		}
// 	}
// 	return table
// }

func (c *Client) normalizeTablePg(table *arrow.Schema, pgTable *arrow.Schema) *arrow.Schema {
	fields := make([]arrow.Field, len(table.Fields()))
	for i, f := range table.Fields() {
		metadata := make(map[string]string, 0)
		if !schema.IsPk(f) {
			metadata[schema.MetadataPrimaryKey] = schema.MetadataFalse
		}
		if c.enabledPks() && schema.IsPk(f) {
			metadata[schema.MetadataPrimaryKey] = schema.MetadataTrue
			f.Nullable = true
		}
		f.Metadata = arrow.MetadataFrom(metadata)
		fields[i] = f
	}
	mdMap := make(map[string]string)
	if pgTable != nil {
		mdMap[schema.MetadataTableName] = schema.TableName(pgTable)
		if constraintName, ok := pgTable.Metadata().GetValue(schema.MetadataConstraintName); ok {
			mdMap[schema.MetadataConstraintName] = constraintName
		}
	}
	mdMap[schema.MetadataTableName] = schema.TableName(table)
	md := arrow.MetadataFrom(mdMap)
	return arrow.NewSchema(fields, &md)
}

func (c *Client) normalizeTable(table *arrow.Schema, pgTable *arrow.Schema) *arrow.Schema {
	switch c.pgType {
	// case pgTypeCockroachDB:
	// 	return c.normalizeTableCockroach(table)
	case pgTypePostgreSQL:
		return c.normalizeTablePg(table, pgTable)
	default:
		panic("unknown pg type")
	}
}

func (c *Client) autoMigrateTable(ctx context.Context, table *arrow.Schema, changes []schema.FieldChange) error {
	tableName := schema.TableName(table)
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if err := c.addColumn(ctx, tableName, change.Current); err != nil {
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

func (*Client) canAutoMigrate(changes []schema.FieldChange) bool {
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if schema.IsPk(change.Current) || !change.Current.Nullable {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if schema.IsPk(change.Previous) || !change.Previous.Nullable {
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
func (c *Client) normalizeTables(tables schema.Schemas, pgTables schema.Schemas) schema.Schemas {
	var result schema.Schemas
	for _, table := range tables {
		pgTabe := pgTables.SchemaByName(schema.TableName(table))
		result = append(result, c.normalizeTable(table, pgTabe))
	}
	return result
}

func (c *Client) nonAutoMigrableTables(tables schema.Schemas, pgTables schema.Schemas) ([]string, [][]schema.FieldChange) {
	var result []string
	var tableChanges [][]schema.FieldChange
	for _, t := range tables {
		pgTable := pgTables.SchemaByName(schema.TableName(t))
		if pgTable == nil {
			continue
		}
		changes := schema.GetSchemaChanges(t, pgTable)
		if !c.canAutoMigrate(changes) {
			result = append(result, schema.TableName(t))
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	tables = c.normalizeTables(tables, pgTables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(tables, pgTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range tables {
		tableName := schema.TableName(table)
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Fields()) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}
		pgTable := pgTables.SchemaByName(tableName)
		if pgTable == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		} else {
			changes := schema.GetSchemaChanges(table, pgTable)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, tableName); err != nil {
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

func (c *Client) addColumn(ctx context.Context, tableName string, column arrow.Field) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Column doesn't exist, creating")
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	sql := "alter table " + tableName + " add column " + columnName + " " + columnType
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *arrow.Schema) error {
	var sb strings.Builder
	tName := schema.TableName(table)
	tableName := pgx.Identifier{tName}.Sanitize()
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(tableName)
	sb.WriteString(" (")
	totalColumns := len(table.Fields())

	primaryKeys := []string{}
	for i, col := range table.Fields() {
		pgType := c.SchemaTypeToPg(col.Type)
		if pgType == "" {
			c.logger.Warn().Str("table", tName).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		columnName := pgx.Identifier{col.Name}.Sanitize()
		fieldDef := columnName + " " + pgType
		if schema.IsUnique(col) {
			fieldDef += " UNIQUE"
		}
		if !col.Nullable {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.enabledPks() && schema.IsPk(col) {
			primaryKeys = append(primaryKeys, col.Name)
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(tName)
		sb.WriteString("_cqpk PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.conn.Exec(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", tableName, err)
	}
	return nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
