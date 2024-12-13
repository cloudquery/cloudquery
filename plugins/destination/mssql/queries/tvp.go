package queries

import (
	"database/sql"
	"reflect"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

type tvpProcQueryBuilder struct {
	Schema string
	Name   string
	Type   string
	Table  *schema.Table
	Values []string
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
	return execTemplate("tvp_drop_proc.sql.tpl", &tvpProcQueryBuilder{Schema: schemaName, Name: procName}),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("procName", procName),
		}
}

func TVPDropType(schemaName string, table *schema.Table) (query string, params []any) {
	typeName := tvpTableType(table)
	return execTemplate("tvp_drop_type.sql.tpl", &tvpProcQueryBuilder{Schema: schemaName, Type: typeName}),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("typeName", typeName),
		}
}

func TVPAddProc(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Schema: schemaName,
		Name:   tvpProcName(table),
		Type:   tvpTableType(table),
		Table:  table,
		Values: GetValueColumns(table),
	}

	return execTemplate("tvp_add_proc.sql.tpl", data)
}

func TVPAddType(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Schema: schemaName,
		Table:  table,
		Type:   tvpTableType(table),
	}

	return execTemplate("tvp_add_type.sql.tpl", data)
}

func TVPQuery(schemaName string, table *schema.Table, records []arrow.Record) (query string, params []any, err error) {
	rows, err := GetRows(array.NewTableFromRecords(table.ToArrowSchema(), records))
	if err != nil {
		return "", nil, err
	}

	return "exec " + sanitizeID(schemaName, tvpProcName(table)) + " @TVP;",
		[]any{
			sql.Named("TVP", mssql.TVP{
				TypeName: sanitizeID(schemaName, tvpTableType(table)),
				Value:    tableTransformer(table)(rows),
			}),
		},
		nil
}

type transformer func([][]any) any

func tableTransformer(table *schema.Table) transformer {
	// 1 prep the fields
	fields := make([]reflect.StructField, len(table.Columns))
	for i, col := range table.Columns {
		fields[i] = reflect.StructField{
			Name: "Fld_" + col.Name,
			Type: reflect.PointerTo(columnGoType(col.Type)), // to account for nullability
		}
	}

	// 2 prep transformer for each field
	row := reflect.StructOf(fields)
	rowSlice := reflect.SliceOf(row)

	rowTransformer := func(rowData []any) reflect.Value {
		v := reflect.New(row).Elem()
		for i, elem := range rowData {
			val := reflect.ValueOf(elem)
			switch {
			case elem == nil, val.IsNil():
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
