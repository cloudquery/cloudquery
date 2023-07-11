package client

import (
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func tablesFromMessages[T message.WriteMessage](messages []T) (schema.Tables, error) {
	tableMap := make(map[string]*schema.Table)
	for _, msg := range messages {
		table := msg.GetTable()
		if _, ok := tableMap[table.Name]; ok {
			continue
		}
		tableMap[table.Name] = table
	}
	tables := make([]*schema.Table, 0, len(tableMap))
	for _, t := range tableMap {
		tables = append(tables, t)
	}
	return tables, nil
}
