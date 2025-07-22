package queries

import (
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func AddColumn(table string, cluster string, col schema.Column) (string, error) {
	definition, err := types.FieldDefinition(col.ToArrowField())
	if err != nil {
		return "", err
	}
	return "ALTER TABLE " + tableNamePart(table, cluster) + " ADD COLUMN IF NOT EXISTS " + definition, nil
}

func DropColumn(table string, cluster string, col schema.Column) string {
	return "ALTER TABLE " + tableNamePart(table, cluster) + " DROP COLUMN IF EXISTS " + util.SanitizeID(col.Name)
}

func SetTTL(table, cluster, ttl string) string {
	if ttl == "" {
		if len(cluster) > 0 {
			return "ALTER TABLE " + util.SanitizeID(table) + " ON CLUSTER " + util.SanitizeID(cluster) + " REMOVE TTL"
		}
		return "ALTER TABLE " + util.SanitizeID(table) + " REMOVE TTL"
	}
	if len(cluster) > 0 {
		return "ALTER TABLE " + util.SanitizeID(table) + " ON CLUSTER " + util.SanitizeID(cluster) + " MODIFY TTL " + ttl
	}
	return "ALTER TABLE " + util.SanitizeID(table) + " MODIFY TTL " + ttl
}
