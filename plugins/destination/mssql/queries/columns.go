package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"golang.org/x/exp/slices"
)

type colQueryBuilder struct {
	Table      string
	Definition *Definition
}

func AddColumn(schemaName string, sc *arrow.Schema, definition *Definition) string {
	return execTemplate("col_add.sql.tpl", &colQueryBuilder{
		Table:      SanitizedTableName(schemaName, sc),
		Definition: definition.sanitized(),
	})
}

func GetPKColumns(sc *arrow.Schema) []string {
	pk := make([]string, 0, len(sc.Fields()))
	for _, field := range sc.Fields() {
		if schema.IsPk(field) {
			pk = append(pk, field.Name)
		}
	}
	return sanitized(slices.Clip(pk)...)
}

func GetValueColumns(sc *arrow.Schema) []string {
	columns := make([]string, 0, len(sc.Fields()))
	for _, field := range sc.Fields() {
		if !schema.IsPk(field) {
			columns = append(columns, field.Name)
		}
	}

	return sanitized(slices.Clip(columns)...)
}

func ColumnNames(sc *arrow.Schema) []string {
	columns := make([]string, len(sc.Fields()))
	for i, field := range sc.Fields() {
		columns[i] = field.Name
	}
	return columns
}
