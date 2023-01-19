package queries

import (
	"database/sql"

	"github.com/cloudquery/plugin-sdk/schema"
)

type readQueryBuilder struct {
	Table              string
	Columns            []string
	CqSourceNameColumn string
	CqSyncTimeColumn   string
}

func Read(schemaName, sourceName string, table *schema.Table) (query string, params []any) {
	return execTemplate("read.sql.tpl",
			&readQueryBuilder{
				Table:              SanitizeID(schemaName, table.Name),
				Columns:            Sanitized(table.Columns.Names()...),
				CqSourceNameColumn: SanitizeID(schema.CqSourceNameColumn.Name),
				CqSyncTimeColumn:   SanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{sql.Named("sourceName", sourceName)}
}
