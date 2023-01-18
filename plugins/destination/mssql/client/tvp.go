package client

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (c *Client) ensureTVP(ctx context.Context, table *schema.Table) error {
	if !c.pkEnabled() {
		return nil
	}

	c.logger.Debug().
		Str("table", table.Name).
		Str("type_name", queries.TVPTableType(c.schemaName, table.Name)).
		Str("proc_name", queries.TVPProcName(c.schemaName, table.Name)).
		Msg("Going tp recreate TVP proc & type")

	_, err := c.db.ExecContext(ctx, queries.TVPDrop(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to drop TVP proc & type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPType(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP type for table %s: %w", table.Name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPProc(c.schemaName, table))
	if err != nil {
		return fmt.Errorf("failed to create TVP proc for table %s: %w", table.Name, err)
	}

	return nil
}

func (c *Client) insertTVP(ctx context.Context, tx *sql.Tx, table *schema.Table, data [][]any) error {
	tf := tableTransformer(table.Columns)

	_, err := tx.ExecContext(ctx, "exec "+queries.TVPProcName(c.schemaName, table.Name)+" @TVP;",
		sql.Named("TVP", mssql.TVP{
			TypeName: queries.TVPTableType(c.schemaName, table.Name),
			Value:    tf(data),
		}),
	)
	return err
}

type transformer func([][]any) any

func tableTransformer(columns schema.ColumnList) transformer {
	// 1 prep the fields
	fld := make([]reflect.StructField, len(columns))
	for i, col := range columns {
		fld[i] = reflect.StructField{
			Name: "Fld_" + col.Name,
			Type: queries.ColumnGoType(col.Type),
		}
	}

	// 2 prep transformer for each field
	row := reflect.StructOf(fld)
	rowSlice := reflect.SliceOf(row)

	rowTransformer := func(rowData []any) reflect.Value {
		v := reflect.New(row).Elem()
		for i, elem := range rowData {
			val := reflect.ValueOf(elem)
			switch {
			case elem == nil:
			case val.IsZero():
			default:
				v.Field(i).Set(val)
			}
		}
		return v
	}

	return func(data [][]any) any {
		rows := reflect.MakeSlice(rowSlice, len(data), len(data))
		for i, elem := range data {
			rows.Index(i).Set(rowTransformer(elem))
		}
		return rows.Interface()
	}
}
