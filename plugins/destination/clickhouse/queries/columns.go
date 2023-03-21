package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AddColumn(table string, cluster string, column *schema.Column) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN " + sanitizeID(column.Name) + " " + chType(column)
}

func DropColumn(table string, cluster string, column *schema.Column) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " DROP COLUMN " + sanitizeID(column.Name)
}
