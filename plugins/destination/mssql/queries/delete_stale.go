package queries

import (
	"database/sql"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type deleteStaleQueryBuilder struct {
	Schema           string
	Table            string
	SourceNameColumn string
	SyncTimeColumn   string
}

func DeleteStale(schemaName string, m *message.WriteDeleteStale) (query string, params []any) {
	return execTemplate("delete_stale.sql.tpl",
			&deleteStaleQueryBuilder{
				Schema:           schemaName,
				Table:            m.TableName,
				SourceNameColumn: sanitizeID(schema.CqSourceNameColumn.Name),
				SyncTimeColumn:   sanitizeID(schema.CqSyncTimeColumn.Name),
			},
		),
		[]any{
			sql.Named("sourceName", m.SourceName),
			sql.Named("syncTime", m.SyncTime),
		}
}
