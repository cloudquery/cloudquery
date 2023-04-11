package queries

import (
	"database/sql"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type deleteStaleQueryBuilder struct {
	Table            string
	SourceNameColumn string
	SyncTimeColumn   string
}

func DeleteStale(schemaName string, table *schema.Table, sourceName string, syncTime time.Time) (query string, params []any) {
	return execTemplate("delete_stale.sql.tpl",
			&deleteStaleQueryBuilder{
				Table:            SanitizedTableName(schemaName, table),
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{
			sql.Named("sourceName", sourceName),
			sql.Named("syncTime", syncTime),
		}
}
