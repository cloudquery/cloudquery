package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
)

func AddColumn(table string, cluster string, field arrow.Field) (string, error) {
	definitions, err := typeconv.FieldDefinitions(field)
	if err != nil {
		return "", err
	}
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN " + definitions[0], nil
}

func DropColumn(table string, cluster string, field arrow.Field) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " DROP COLUMN " + util.SanitizeID(field.Name)
}

func ColumnNames(fields []arrow.Field) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = field.Name
	}
	return res
}
