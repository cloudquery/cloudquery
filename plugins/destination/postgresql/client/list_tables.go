package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// this returns the following table in sorted manner:
// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
// | ordinal_position | table_name | column_name | data_type | is_primary_key| not_null  | pk_constraint_name  |
// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
// |              1 | users       | id          | bigint     | YES           | true 		 | cq_users_pk 	  	   |
// |              2 | users       | name        | text       | NO            | false 	   | 					           |
// |              3 | users       | email       | text       | NO            | true 		 | cq_users_pk         |
// |              1 | posts       | id          | bigint     | YES           | true 		 | cq_posts_pk			   |
// |              2 | posts       | title       | text       | NO            | false 	   | 					           |
const selectTables = `
SELECT
	columns.ordinal_position AS ordinal_position,
	pg_class.relname AS table_name,
	pg_attribute.attname AS column_name,
	CASE
	    -- This is required per the differences in pg_catalog.format_type implementations
	    -- between PostgreSQL & CockroachDB.
	    -- namely, numeric(20,0)[] is returned as numeric[] unless we use the typelem format + []
	    WHEN pg_type.typcategory = 'A' AND pg_type.typelem != 0
		THEN pg_catalog.format_type(pg_type.typelem, pg_attribute.atttypmod) || '[]'
		ELSE pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod)
	END AS data_type,
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
	pg_catalog.pg_type ON pg_type.oid = pg_attribute.atttypid
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
	%s 
ORDER BY
	table_name ASC, ordinal_position ASC;
`

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	c.pgTablesToPKConstraints = map[string]string{}
	var tables schema.Tables
	var whereClause string
	if c.pgType == pgTypeCockroachDB {
		whereClause = " AND information_schema.columns.is_hidden != 'YES'"
	}
	q := fmt.Sprintf(selectTables, c.currentSchemaName, whereClause)
	rows, err := c.conn.Query(ctx, q)
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
		if ordinalPosition == 1 {
			tables = append(tables, &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			})
		}
		table := tables[len(tables)-1]
		if pkName != "" {
			c.pgTablesToPKConstraints[tableName], table.PkConstraintName = pkName, pkName
		}
		table.Columns = append(table.Columns, schema.Column{
			Name:       columnName,
			PrimaryKey: isPrimaryKey,
			NotNull:    notNull,
			Type:       c.PgToSchemaType(columnType),
		})
	}
	return tables, nil
}
