package queries

import (
	"database/sql"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type readQueryBuilder struct {
	Table            string
	Columns          []string
	SourceNameColumn string
	SyncTimeColumn   string
}

func Read(schemaName, sourceName string, table *schema.Table) (query string, params []any) {
	return execTemplate("read.sql.tpl",
			&readQueryBuilder{
				Table:            SanitizedTableName(schemaName, table),
				Columns:          sanitized(table.Columns.Names()...),
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{sql.Named("sourceName", sourceName)}
}
