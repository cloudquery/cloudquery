package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
)

func AddColumn(table string, cluster string, field arrow.Field) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN " + sanitizeID(field.Name) + " " + chFieldType(field)
}

func DropColumn(table string, cluster string, field arrow.Field) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " DROP COLUMN " + sanitizeID(field.Name)
}

func ColumnNames(fields []arrow.Field) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = field.Name
	}
	return res
}
