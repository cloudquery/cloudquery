package client

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// this returns the following table in sorted manner:
// +----------------+-------------+-------------+------------+----------------+-----------+-----------+---------------------+
// | ordinal_position | table_name | column_name | data_type | is_primary_key | not_null  | is_unique | constraint_name  |
// +----------------+-------------+-------------+------------+----------------+-----------+-----------+---------------------+
// |              1 | users       | id          | bigint     | YES            | true      | true      | cq_users_pk         |
// |              2 | users       | name        | text       | NO             | false     | false     |                     |
// |              3 | users       | email       | text       | NO             | true      | false     | cq_users_pk         |
// |              1 | posts       | id          | bigint     | YES            | true      | true      | cq_posts_pk         |
// |              2 | posts       | title       | text       | NO             | false     | false     |                     |
//
//go:embed sql/list_tables.sql
var selectTables string

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	var tables schema.Tables
	q := fmt.Sprintf(selectTables, c.currentSchemaName)
	rows, err := c.Conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tableMap := make(map[string]*schema.Table)

	for rows.Next() {
		var ordinalPosition int
		var tableName, columnName, columnType, constraintName string
		var isPrimaryKey, notNull, isUnique bool
		if err := rows.Scan(&ordinalPosition, &tableName, &columnName, &columnType, &isPrimaryKey, &notNull, &isUnique, &constraintName); err != nil {
			return nil, err
		}
		table := tableMap[tableName]
		if table == nil {
			table = &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			}
			tableMap[tableName] = table
			tables = append(tables, table)
		}

		if isPrimaryKey && constraintName != "" {
			table.PkConstraintName = constraintName
		}
		table.Columns = append(table.Columns, schema.Column{
			Name:       columnName,
			PrimaryKey: isPrimaryKey,
			NotNull:    notNull,
			Unique:     isUnique,
			Type:       c.PgToSchemaType(columnType),
		})
	}
	return tables, nil
}
