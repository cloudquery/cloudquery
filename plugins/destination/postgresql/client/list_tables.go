package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
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
				Name: tableName,
			})
		}
		table := tables[len(tables)-1]
		if pkName != "" {
			table.PkConstraintName = pkName
		}
		schemaType := c.PgToSchemaType(columnType)
		table.Columns = append(table.Columns, schema.Column{
			Name:       columnName,
			Type:       schemaType,
			PrimaryKey: isPrimaryKey,
			NotNull:    notNull,
		})
	}
	return tables, nil
}
