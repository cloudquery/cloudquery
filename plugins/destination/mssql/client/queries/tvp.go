package queries

import (
	"database/sql"
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

type tvpProcQueryBuilder struct {
	Name    string
	Type    string
	Table   string
	Columns Definitions
	PK      []string
	Values  []string
}

func tvpProcName(schemaName string, tableName string) string {
	const pfx = "cq_proc_"
	return sanitizeID(schemaName, pfx+tableName)
}

func tvpTableType(schemaName string, tableName string) string {
	const pfx = "cq_tbl_"
	return sanitizeID(schemaName, pfx+tableName)
}

func TVPDrop(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Name: tvpProcName(schemaName, table.Name),
		Type: tvpTableType(schemaName, table.Name),
	}

	return execTemplate("tvp_drop.sql.tpl", data)
}

func TVPProc(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Name:   tvpProcName(schemaName, table.Name),
		Type:   tvpTableType(schemaName, table.Name),
		Table:  sanitizeID(schemaName, table.Name),
		PK:     GetPKColumns(table, true),
		Values: GetValueColumns(table.Columns),
	}

	return execTemplate("tvp_proc.sql.tpl", data)
}

func TVPType(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Type:    tvpTableType(schemaName, table.Name),
		Columns: GetDefinitions(table.Columns, true),
	}

	return execTemplate("tvp_type.sql.tpl", data)
}

func TVPQuery(schemaName string, table *schema.Table, data [][]any) (query string, params []any) {
	tf := tableTransformer(table.Columns)

	return "exec " + tvpProcName(schemaName, table.Name) + " @TVP;",
		[]any{
			sql.Named("TVP", mssql.TVP{
				TypeName: tvpTableType(schemaName, table.Name),
				Value:    tf(data),
			}),
		}
}

type transformer func([][]any) any

func tableTransformer(columns schema.ColumnList) transformer {
	// 1 prep the fields
	fld := make([]reflect.StructField, len(columns))
	for i, col := range columns {
		fld[i] = reflect.StructField{
			Name: "Fld_" + col.Name,
			Type: columnGoType(col.Type),
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
