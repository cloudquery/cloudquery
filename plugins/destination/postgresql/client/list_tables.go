package client

import (
	"context"
	"fmt"
	"strings"

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
	%s 
ORDER BY
	table_name ASC, ordinal_position ASC;
`

func (c *Client) listTables(ctx context.Context, include, exclude []string) (schema.Tables, error) {
	var tables schema.Tables
	whereClause := c.whereClause(include, exclude)
	if c.pgType == pgTypeCockroachDB {
		whereClause += " AND information_schema.columns.is_hidden != 'YES'"
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
			table.PkConstraintName = pkName
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

func (c *Client) whereClause(include, exclude []string) string {
	if len(include) == 0 && len(exclude) == 0 {
		return ""
	}
	var where string
	if len(include) > 0 {
		where = fmt.Sprintf("AND pg_class.relname IN (%s)", c.inClause(include))
	}
	if len(exclude) > 0 {
		where = fmt.Sprintf("AND pg_class.relname NOT IN (%s)", c.inClause(exclude))
	}
	return where
}

func (*Client) inClause(values []string) string {
	var inClause string
	for i, value := range values {
		value = strings.ReplaceAll(value, "'", "")  // strip single quotes
		value = strings.ReplaceAll(value, "*", "%") // replace * with %
		if i == 0 {
			inClause = fmt.Sprintf("'%s'", value)
			continue
		}
		inClause += fmt.Sprintf(", '%s'", value)
	}
	return inClause
}
