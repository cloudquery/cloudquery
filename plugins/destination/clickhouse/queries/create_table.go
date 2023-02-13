package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type createTableQueryBuilder struct {
	Table      *schema.Table
	SortingKey []string
}

func CreateTable(table *schema.Table) string {
	return execTemplate("create_table.sql.tpl", &createTableQueryBuilder{
		Table:      normalizeTable(table),
		SortingKey: []string{schema.CqIDColumn.Name}, // only _cq_id for now
	})
}
