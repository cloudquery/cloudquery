package queries

import (
	"database/sql"
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
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

func tvpProcName(sc *arrow.Schema) string {
	const pfx = "cq_proc_"
	return pfx + schema.TableName(sc)
}

func tvpTableType(sc *arrow.Schema) string {
	const pfx = "cq_tbl_"
	return pfx + schema.TableName(sc)
}

func TVPDropProc(schemaName string, sc *arrow.Schema) (query string, params []any) {
	procName := tvpProcName(sc)
	data := &tvpProcQueryBuilder{
		Name: sanitizeID(schemaName, procName),
	}

	return execTemplate("tvp_drop_proc.sql.tpl", data),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("procName", procName),
		}
}

func TVPDropType(schemaName string, sc *arrow.Schema) (query string, params []any) {
	typeName := tvpTableType(sc)
	data := &tvpProcQueryBuilder{
		Type: sanitizeID(schemaName, typeName),
	}

	return execTemplate("tvp_drop_type.sql.tpl", data),
		[]any{
			sql.Named("schemaName", schemaName),
			sql.Named("typeName", typeName),
		}
}

func TVPAddProc(schemaName string, sc *arrow.Schema) string {
	data := &tvpProcQueryBuilder{
		Name:        sanitizeID(schemaName, tvpProcName(sc)),
		Type:        sanitizeID(schemaName, tvpTableType(sc)),
		Table:       sanitizeID(schemaName, schema.TableName(sc)),
		PK:          GetPKColumns(sc),
		Values:      GetValueColumns(sc),
		ColumnNames: sanitized(ColumnNames(sc)...),
	}

	return execTemplate("tvp_add_proc.sql.tpl", data)
}

func TVPAddType(schemaName string, sc *arrow.Schema) string {
	data := &tvpProcQueryBuilder{
		Type:    sanitizeID(schemaName, tvpTableType(sc)),
		Columns: GetDefinitions(sc, true),
	}

	return execTemplate("tvp_add_type.sql.tpl", data)
}

func TVPQuery(schemaName string, sc *arrow.Schema, records []arrow.Record) (query string, params []any, err error) {
	reader, err := array.NewRecordReader(sc, records)
	if err != nil {
		return "", nil, err
	}

	rows, err := GetRows(reader)
	if err != nil {
		return "", nil, err
	}

	return "exec " + sanitizeID(schemaName, tvpProcName(sc)) + " @TVP;",
		[]any{
			sql.Named("TVP", mssql.TVP{
				TypeName: sanitizeID(schemaName, tvpTableType(sc)),
				Value:    tableTransformer(sc)(rows),
			}),
		},
		nil
}

type transformer func([][]any) any

func tableTransformer(sc *arrow.Schema) transformer {
	// 1 prep the fields
	fld := make([]reflect.StructField, len(sc.Fields()))
	for i, field := range sc.Fields() {
		fld[i] = reflect.StructField{
			Name: "Fld_" + field.Name,
			Type: reflect.PointerTo(columnGoType(field.Type)), // to account for nullability
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
