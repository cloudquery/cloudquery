package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func (c *Client) pkEnabled() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}

func (c *Client) getTableColumns(ctx context.Context, table *schema.Table) (schema.ColumnList, error) {
	query, params := queries.GetTableSchema(c.schemaName, table)
	var tc schema.ColumnList

	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var typ string
		var nullable string
		var charMaxLength *string

		if err := row.Scan(&name, &typ, &nullable, &charMaxLength); err != nil {
			return err
		}

		if (typ == "nvarchar" || typ == "varbinary") && charMaxLength != nil {
			if *charMaxLength == "-1" {
				*charMaxLength = "max"
			}
			typ += "(" + *charMaxLength + ")"
		}

		if typ == "datetimeoffset" {
			return fmt.Errorf(`column %q from table %q is of type "datetimeoffset" which was changed to "datetime2". Please drop the database to upgrade to this version`, name, table.Name)
		}

		schemaType, err := queries.SchemaType(table.Name, name, typ)
		if err != nil {
			return err
		}
		tc = append(tc, schema.Column{Name: name, Type: schemaType, CreationOptions: schema.ColumnCreationOptions{NotNull: nullable == "NO"}})

		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	return tc, nil
}

func (c *Client) getTablePK(ctx context.Context, table *schema.Table) ([]string, error) {
	query, params := queries.GetTablePK(c.schemaName, table)

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
