package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
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
		 WHEN contype = 'p' THEN true
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
 AND pg_constraint.conkey[1] = pg_attribute.attnum
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
		// This is a small optimization to hold only the tables that are in the plugin
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
				NotNull: 	notNull,
			},
			Resolver: func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
				return resource.Set(columnName, resource.Item.([]interface{})[ordinalPosition-1])
			},
			Type: c.PgToSchemaType(columnType),
		})
	}
	return tables, nil
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables, options destination.MigrateOptions) error {
	options.Force = true
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	for _, table := range tables.FlattenTables() {
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
		}

		pgTable := pgTables.Get(table.Name)
		if pgTable != nil {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, pgTable, table, options); err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) alterColumn(ctx context.Context, tableName string, column schema.Column) error {
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	sql := "alter table " + tableName + " drop column " + columnName
	// right now we will drop the column and re-create. in the future we will have an option to automigrate
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop column %s on table %s: %w", column.Name, tableName, err)
	}
	sql = "alter table " + tableName + " add column " + columnName + " " + columnType
	if column.CreationOptions.NotNull {
		sql += " not null"
	}
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
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
	dropConstraintSQL := "alter table " + tableName + " drop constraint if exists " + pgTable.PkConstraintName
	if _, err := tx.Exec(ctx, dropConstraintSQL); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			c.logger.Error().Err(err).Msg("failed to rollback transaction")
		}
		return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
	}
	
	sql := "alter table " + tableName + " add constraint primary key (" + strings.Join(table.PrimaryKeys(), ",") + ")"
	if _, err := tx.Exec(ctx, sql); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			c.logger.Error().Err(err).Msg("failed to rollback transaction")
		}
		return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction to recreate primary keys: %w", err)
	}
	return nil
}

func (c *Client) autoMigrateTable(ctx context.Context, pgTable *schema.Table, table *schema.Table, options destination.MigrateOptions) error {
	if changedColumns := table.GetChangedColumns(pgTable); changedColumns != nil {
		if !options.Force {
			return fmt.Errorf("postgres table %s has changed %v columns, use --force to drop the column", table.Name, changedColumns)
		}
		for _, col := range changedColumns {
			if err := c.alterColumn(ctx, table.Name, col); err != nil {
				return err
			}
		}
	}

	if addedColumns := table.GetAddedColumns(pgTable); addedColumns != nil {
		for _, col := range addedColumns {
			if err := c.addColumn(ctx, table.Name, col); err != nil {
				return err
			}
		}
	}

	if c.enabledPks() && !table.IsPkEqual(pgTable) {
		if !options.Force {
			return fmt.Errorf("postgres table %s primary key is different from the schema, use --force to drop the constraint", table.Name)
		}
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

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
