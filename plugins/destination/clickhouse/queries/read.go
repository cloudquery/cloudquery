package queries

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/plugin-sdk/schema"
)

type readQueryBuilder struct {
	Table            string
	Columns          []string
	SourceNameColumn string
	SyncTimeColumn   string
}

func Read(sourceName string, table *schema.Table) (query string, params []any) {
	return execTemplate("read.sql.tpl",
			&readQueryBuilder{
				Table:            table.Name,
				Columns:          table.Columns.Names(),
				SourceNameColumn: schema.CqSourceNameColumn.Name,
				SyncTimeColumn:   schema.CqSyncTimeColumn.Name,
			},
		),
		[]any{driver.NamedValue{Name: "sourceName", Value: sourceName}}
}
