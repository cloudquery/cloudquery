package client

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5"
)

// this returns the following table in sorted manner:
// +----------------+-------------+-------------+------------+--------------+
// | ordinal_position | table_name | column_name | data_type | is_primary_key|
// +----------------+-------------+-------------+------------+--------------+
// |              1 | users       | id          | bigint     | YES          |
// |              2 | users       | name        | text       | NO           |
// |              3 | users       | email       | text       | NO           |
// |              1 | posts       | id          | bigint     | YES          |
// |              2 | posts       | title       | text       | NO           |
const selectTables = `
SELECT
columns.ordinal_position AS ordinal_position,
pg_class.relname AS table_name,
pg_attribute.attname AS column_name,
pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod) AS data_type,
CASE 
    WHEN contype = 'p' THEN 'YES'
    ELSE 'NO'
END AS is_primary_key
FROM
pg_catalog.pg_attribute
INNER JOIN
pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
INNER JOIN
pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
LEFT JOIN
pg_catalog.pg_constraint ON pg_constraint.conrelid = pg_attribute.attrelid
AND pg_constraint.conkey[1] = pg_attribute.attnum
INNER JOIN
information_schema.columns ON columns.table_name = pg_class.relname AND columns.column_name = pg_attribute.attname AND columns.table_schema = pg_catalog.pg_namespace.nspname
WHERE
pg_attribute.attnum > 0
AND NOT pg_attribute.attisdropped
AND pg_catalog.pg_namespace.nspname in (SELECT TRIM(BOTH from unnest(string_to_array(reset_val,','))) FROM pg_settings WHERE name='search_path')
ORDER BY
table_name ASC , ordinal_position ASC;
`

func (c *Client) ListTables(ctx context.Context) (schema.Tables, error) {
	var tables schema.Tables
	rows, err := c.Conn.Query(ctx, selectTables)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ordinalPosition int
		var isPrimaryKey, tableName, columnName, columnType string
		if err := rows.Scan(&ordinalPosition, &tableName, &columnName, &columnType, &isPrimaryKey); err != nil {
			return nil, err
		}
		if ordinalPosition == 1 {
			tables = append(tables, &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			})
		}
		table := tables[len(tables)-1]
		primaryKey := false
		if isPrimaryKey == "YES" {
			primaryKey = true
		}
		table.Columns = append(table.Columns, schema.Column{
			Name: columnName,
			CreationOptions: schema.ColumnCreationOptions{
				PrimaryKey: primaryKey,
			},
			Resolver: func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
				return resource.Set(columnName, resource.Item.([]interface{})[ordinalPosition-1])
			},
			Type: c.PgToSchemaType(columnType),
		})
	}
	for i := range tables {
		tables[i].Resolver = createTableResolver(tables[i])
	}
	return tables, nil
}

func createTableResolver(table *schema.Table) schema.TableResolver {
	colNames := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		colNames[i] = pgx.Identifier{col.Name}.Sanitize()
	}
	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + pgx.Identifier{table.Name}.Sanitize()
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		c := meta.(*Client)
		rows, err := c.Conn.Query(ctx, query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}
			res <- []interface{}{values}
		}
		return nil
	}
}
