package queries

import (
	"database/sql"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type deleteStaleQueryBuilder struct {
	Schema           string
	Table            string
	SourceNameColumn string
	SyncTimeColumn   string
}

func DeleteStale(schemaName string, tableName, sourceName string, syncTime time.Time) (query string, params []any) {
	return execTemplate("delete_stale.sql.tpl",
			&deleteStaleQueryBuilder{
				Schema:           schemaName,
				Table:            tableName,
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{
			sql.Named("sourceName", sourceName),
			sql.Named("syncTime", syncTime),
		}
}
