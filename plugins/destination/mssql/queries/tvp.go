package queries

import (
	"database/sql"
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

type tvpProcQueryBuilder struct {
	Name        string
	Type        string
	Table       string
	Columns     Definitions
	ColumnNames []string
	PK          []string
	Values      []string
}

func tvpProcName(table *schema.Table) string {
	const pfx = "cq_proc_"
	return pfx + table.Name
}

func tvpTableType(table *schema.Table) string {
	const pfx = "cq_tbl_"
	return pfx + table.Name
}

func TVPDropProc(schemaName string, table *schema.Table) (query string, params []any) {
	procName := tvpProcName(table)
	data := &tvpProcQueryBuilder{
		Name: sanitizeID(schemaName, procName),
	}

	return execTemplate("tvp_drop_proc.sql.tpl", data),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("procName", procName),
		}
}

func TVPDropType(schemaName string, table *schema.Table) (query string, params []any) {
	typeName := tvpTableType(table)
	data := &tvpProcQueryBuilder{
		Type: sanitizeID(schemaName, typeName),
	}

	return execTemplate("tvp_drop_type.sql.tpl", data),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("typeName", typeName),
		}
}

func TVPAddProc(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Name:        sanitizeID(schemaName, tvpProcName(table)),
		Type:        sanitizeID(schemaName, tvpTableType(table)),
		Table:       sanitizeID(schemaName, table.Name),
		PK:          GetPKColumns(table),
		Values:      GetValueColumns(table.Columns),
		ColumnNames: sanitized(table.Columns.Names()...),
	}

	return execTemplate("tvp_add_proc.sql.tpl", data)
}

func TVPAddType(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Type:    sanitizeID(schemaName, tvpTableType(table)),
		Columns: GetDefinitions(table.Columns, true),
	}

	return execTemplate("tvp_add_type.sql.tpl", data)
}

func TVPQuery(schemaName string, table *schema.Table, data [][]any) (query string, params []any) {
	tf := tableTransformer(table.Columns)

	return "exec " + sanitizeID(schemaName, tvpProcName(table)) + " @TVP;",
		[]any{
			sql.Named("TVP", mssql.TVP{
				TypeName: sanitizeID(schemaName, tvpTableType(table)),
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
