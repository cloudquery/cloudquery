package queries

import (
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func AddColumn(table string, cluster string, col schema.Column) (string, error) {
	definition, err := types.FieldDefinition(col.ToArrowField())
	if err != nil {
		return "", err
	}
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN " + definition, nil
}

func DropColumn(table string, cluster string, col schema.Column) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " DROP COLUMN " + util.SanitizeID(col.Name)
}
