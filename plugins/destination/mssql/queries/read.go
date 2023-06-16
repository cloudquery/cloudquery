package queries

import (
	"database/sql"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

type readQueryBuilder struct {
	Schema           string
	Table            *schema.Table
	SourceNameColumn string
	SyncTimeColumn   string
}

func Read(schemaName, sourceName string, table *schema.Table) (query string, params []any) {
	return execTemplate("read.sql.tpl",
			&readQueryBuilder{
				Schema:           schemaName,
				Table:            table,
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{sql.Named("sourceName", sourceName)}
}
