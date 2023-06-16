package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/exp/slices"
)

func (c *Client) pkEnabled() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}

func (c *Client) getTableColumns(ctx context.Context, tableName string, pks []string) (schema.ColumnList, error) {
	query, params := queries.GetTableSchema(c.schemaName, tableName)

	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	columns := make(schema.ColumnList, 0)
	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var sqlType string
		var nullable string
		var charMaxLength *string

		if err := row.Scan(&name, &sqlType, &nullable, &charMaxLength); err != nil {
			return err
		}

		if (sqlType == "nvarchar" || sqlType == "varbinary") && charMaxLength != nil {
			if *charMaxLength == "-1" {
				*charMaxLength = "max"
			}
			sqlType += "(" + *charMaxLength + ")"
		}

		dataType := queries.SchemaType(sqlType)

		columns = append(columns, schema.Column{
			Name:       name,
			Type:       dataType,
			PrimaryKey: slices.Contains(pks, name),
			NotNull:    nullable == "NO",
		})

		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	return slices.Clip(columns), nil
}

func (c *Client) getTablePK(ctx context.Context, tableName string) ([]string, error) {
	query, params := queries.GetTablePK(c.schemaName, tableName)

	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	var result []string
	if err := processRows(rows, func(row *sql.Rows) error {
		var name string

		if err := rows.Scan(&name); err != nil {
			return err
		}

		result = append(result, name)

		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	return result, nil
}
