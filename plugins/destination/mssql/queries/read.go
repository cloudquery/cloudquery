package queries

import (
	"database/sql"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type readQueryBuilder struct {
	Table            string
	Columns          []string
	SourceNameColumn string
	SyncTimeColumn   string
}

func Read(schemaName, sourceName string, sc *arrow.Schema) (query string, params []any) {
	return execTemplate("read.sql.tpl",
			&readQueryBuilder{
				Table:            SanitizedTableName(schemaName, sc),
				Columns:          sanitized(ColumnNames(sc)...),
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{sql.Named("sourceName", sourceName)}
}
