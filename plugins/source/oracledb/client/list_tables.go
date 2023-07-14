package client

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type column struct {
	id       int
	name     string
	dataType string
	notNull  bool
}

func Identifier(name string) string {
	return "\"" + name + "\""
}

func (c *Client) updateTableConstraints(ctx context.Context, table *schema.Table) error {
	query := `SELECT CONSTRAINT_TYPE, COLUMN_NAME FROM USER_CONSTRAINTS NATURAL JOIN USER_CONS_COLUMNS WHERE TABLE_NAME = :1`
	rows, err := c.db.QueryContext(ctx, query, table.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	constraints := make(map[string][]string)
	for rows.Next() {
		var constraintType string
		var columnName string

		if err := rows.Scan(&constraintType, &columnName); err != nil {
			return err
		}
		constraints[columnName] = append(constraints[columnName], constraintType)
	}

	for i, column := range table.Columns {
		if constraintTypes, ok := constraints[column.Name]; ok {
			for _, constraintType := range constraintTypes {
				switch constraintType {
				case "P":
					table.Columns[i].PrimaryKey = true
				case "U":
					table.Columns[i].Unique = true
				}
			}
		}
	}

	return nil
}

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	// Please note we don't use ORDER BY here because it's slower than sorting in memory via Go sort.SliceStable
	// See more about the default namespaces we exclude in https://www.oracletutorial.com/oracle-administration/oracle-tablespace/
	// Also invisible columns have COLUMN_ID = NULL so we exclude them as well. See https://forums.oracle.com/ords/apexds/post/oracle-12c-invisible-columns-0462
	query := `
	SELECT ALL_TAB_COLS.TABLE_NAME, COLUMN_ID, COLUMN_NAME, DATA_TYPE, DATA_LENGTH, NULLABLE, DATA_PRECISION, DATA_SCALE FROM ALL_TABLES
	LEFT JOIN ALL_TAB_COLS ON ALL_TABLES.TABLE_NAME = ALL_TAB_COLS.TABLE_NAME
	WHERE ALL_TABLES.TABLESPACE_NAME NOT IN ('SYSAUX', 'SYSTEM', 'TEMP', 'UNDOTBS1')
	AND ALL_TAB_COLS.COLUMN_ID IS NOT NULL
	`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schemaTables := make(map[string][]column)
	for rows.Next() {
		var tableName string
		var columnId int
		var columnName string
		var dataType string
		var dataLength int
		var nullable string
		var precision sql.NullInt64
		var scale sql.NullInt64

		if err := rows.Scan(&tableName, &columnId, &columnName, &dataType, &dataLength, &nullable, &precision, &scale); err != nil {
			return nil, err
		}
		dataType = strings.ToLower(dataType)
		if dataType == "char" || dataType == "raw" {
			dataType = fmt.Sprintf("%s(%d)", dataType, dataLength)
		}
		if dataType == "number" {
			if precision.Valid && scale.Valid {
				dataType = fmt.Sprintf("number(%d,%d)", precision.Int64, scale.Int64)
			} else if precision.Valid {
				dataType = fmt.Sprintf("number(%d)", precision.Int64)
			}
		}
		schemaTables[tableName] = append(schemaTables[tableName], column{
			id:       columnId,
			name:     columnName,
			dataType: dataType,
			notNull:  nullable == "N",
		})
	}

	tables := make(schema.Tables, 0, len(schemaTables))
	for tableName, columns := range schemaTables {
		table := schema.Table{
			Name: tableName,
		}
		sort.SliceStable(columns, func(i, j int) bool {
			return columns[i].id < columns[j].id
		})
		for _, column := range columns {
			table.Columns = append(table.Columns, schema.Column{
				Name:    column.name,
				Type:    SchemaType(column.dataType),
				NotNull: column.notNull,
			})
		}
		err := c.updateTableConstraints(ctx, &table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}

	sort.SliceStable(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})

	return tables, nil
}
