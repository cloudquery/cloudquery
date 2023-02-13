package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type colQueryBuilder struct {
	Table  string
	Column *schema.Column
}

func AddColumn(table string, column *schema.Column) string {
	return execTemplate("col_add.sql.tpl", &colQueryBuilder{Table: table, Column: column})
}

func DropColumn(table string, column *schema.Column) string {
	return execTemplate("col_drop.sql.tpl", &colQueryBuilder{Table: table, Column: column})
}
