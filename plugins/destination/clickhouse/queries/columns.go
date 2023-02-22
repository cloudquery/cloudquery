package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AddColumn(table string, column *schema.Column) string {
	return "ALTER TABLE " + sanitizeID(table) + " ADD COLUMN " + sanitizeID(column.Name) + " " + chType(column)
}

func DropColumn(table string, column *schema.Column) string {
	return "ALTER TABLE " + sanitizeID(table) + " DROP COLUMN " + sanitizeID(column.Name)
}
