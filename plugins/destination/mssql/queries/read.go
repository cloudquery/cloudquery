package queries

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type readQueryBuilder struct {
	Schema string
	Table  *schema.Table
}

func Read(schemaName string, table *schema.Table) string {
	return execTemplate("read.sql.tpl",
		&readQueryBuilder{Schema: schemaName, Table: table},
	)
}
