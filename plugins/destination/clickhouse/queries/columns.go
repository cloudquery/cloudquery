package queries

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
)

func AddColumn(table string, cluster string, field arrow.Field) (string, error) {
	definition, err := types.FieldDefinition(field)
	if err != nil {
		return "", err
	}
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN " + definition, nil
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
