package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type (
	createTableQueryBuilder struct {
		Table      *TableDefinition
		SortingKey []string
	}
)

func CreateTable(table *TableDefinition) string {
	return execTemplate("create_table.sql.tpl", &createTableQueryBuilder{
		Table:      table,
		SortingKey: []string{schema.CqIDColumn.Name}, // only _cq_id for now
	})
}
