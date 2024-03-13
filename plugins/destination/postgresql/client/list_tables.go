package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// this returns the following table in sorted manner:
// +----------------+-------------+-------------+------------+---------------+-----------+----------------------+-----------+---------------------------+
// | ordinal_position | table_name | column_name | data_type | is_primary_key| not_null  | pk_constraint_name   | is_unique | unique_constraint_name    |
// +----------------+-------------+-------------+------------+---------------+-----------+----------------------+-----------+---------------------------+
// |              1 | users       | id          | bigint     | YES           | true 		 | cq_users_pk		| false		|							|
// |              2 | users       | name        | text       | NO            | false 	     |					| true		| cq_users_name_unique		|
// |              3 | users       | email       | text       | NO            | true 		 | cq_users_pk		| false		|							|
// |              1 | posts       | id          | bigint     | YES           | true 		 | cq_posts_pk		| false		|							|
// |              2 | posts       | title       | text       | NO            | false 	     |					| false		|							|
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
		WHEN pk_constraint.conkey IS NOT NULL AND array_position(pk_constraint.conkey, pg_attribute.attnum) > 0 THEN true
		ELSE false
	END AS is_primary_key,
	CASE 
		WHEN pg_attribute.attnotnull THEN true
		ELSE false
	END AS not_null,
	CASE 
		WHEN pk_constraint.conkey IS NOT NULL AND array_position(pk_constraint.conkey, pg_attribute.attnum) > 0 THEN COALESCE(pk_constraint.conname, '')
		ELSE ''
	END AS primary_key_constraint_name,
	CASE 
		WHEN unique_constraint.conkey IS NOT NULL AND array_position(unique_constraint.conkey, pg_attribute.attnum) > 0 THEN true
		ELSE false
	END AS is_unique,
	CASE 
		WHEN unique_constraint.conkey IS NOT NULL AND array_position(unique_constraint.conkey, pg_attribute.attnum) > 0 THEN COALESCE(unique_constraint.conname, '')
		ELSE ''
	END AS unique_constraint_name
FROM
	pg_catalog.pg_attribute
	INNER JOIN
	pg_catalog.pg_type ON pg_type.oid = pg_attribute.atttypid
	INNER JOIN
	pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
	INNER JOIN
	pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
	LEFT JOIN
	pg_catalog.pg_constraint as pk_constraint ON pk_constraint.conrelid = pg_attribute.attrelid
	AND pk_constraint.conkey IS NOT NULL AND pk_constraint.contype = 'p' AND array_position(pk_constraint.conkey, pg_attribute.attnum) > 0
	LEFT JOIN
	pg_catalog.pg_constraint unique_constraint ON unique_constraint.conrelid = pg_attribute.attrelid
	AND unique_constraint.conkey IS NOT NULL AND unique_constraint.contype = 'u' AND array_position(unique_constraint.conkey, pg_attribute.attnum) > 0

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
	c.pgTablesToPKConstraints = map[string]*pkConstraintDetails{}
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
		var tableName, columnName, columnType, pkName, uniqueName string
		var isPrimaryKey, notNull, isUnique bool
		if err := rows.Scan(&ordinalPosition, &tableName, &columnName, &columnType, &isPrimaryKey, &notNull, &pkName, &isUnique, &uniqueName); err != nil {
			return nil, err
		}
		if ordinalPosition == 1 {
			tables = append(tables, &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			})
		}
		table := tables[len(tables)-1]

		// We always want to record that we saw the table, even if it doesn't have a PK constraint.
		entry, ok := c.pgTablesToPKConstraints[tableName]
		if !ok {
			entry = new(pkConstraintDetails)
			c.pgTablesToPKConstraints[tableName] = entry
		}
		if pkName != "" {
			entry.name = pkName
			// we want to store any current column that is part of the PK
			entry.columns = append(entry.columns, columnName)
			table.PkConstraintName = pkName
		}

		table.Columns = append(table.Columns, schema.Column{
			Name:       columnName,
			PrimaryKey: isPrimaryKey,
			NotNull:    notNull,
			Type:       c.PgToSchemaType(columnType),
			// We check that the unique constraint name is the default single column name
			// This will ensure we don't get mixed up with multi-column unique constraints
			// which we don't support.
			Unique: isUnique && uniqueName == tableName+"_"+columnName+"_key",
		})
	}
	return tables, nil
}
