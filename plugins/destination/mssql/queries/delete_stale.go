package queries

import (
	"database/sql"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type deleteStaleQueryBuilder struct {
	Table            string
	SourceNameColumn string
	SyncTimeColumn   string
}

func DeleteStale(schemaName string, sc *arrow.Schema, sourceName string, syncTime time.Time) (query string, params []any) {
	return execTemplate("delete_stale.sql.tpl",
			&deleteStaleQueryBuilder{
				Table:            SanitizedTableName(schemaName, sc),
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{
			sql.Named("sourceName", sourceName),
			sql.Named("syncTime", syncTime),
		}
}
