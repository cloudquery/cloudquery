package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	for r := range res {
		switch m := r.(type) {
		case *message.WriteMigrateTable:
			if err := c.migrate(ctx, m.MigrateForce, schema.Tables{m.Table}); err != nil {
				return fmt.Errorf("failed to process MigrateTable message: %w", err)
			}
		case *message.WriteDeleteStale:
			if err := c.deleteStale(ctx, m.TableName, m.SourceName, m.SyncTime); err != nil {
				return fmt.Errorf("failed to process DeleteStale message: %w", err)
			}
		case *message.WriteInsert:
			if err := c.insertMessage(ctx, m); err != nil {
				return fmt.Errorf("failed to process Insert message: %w", err)
			}
		default:
			return fmt.Errorf("unsupported message type: %T", m)
		}
	}
	return nil
}

func (c *Client) insertMessage(ctx context.Context, m *message.WriteInsert) error {
	table := m.GetTable()
	sc := m.Record.Schema()
	var sql string
	if len(table.PrimaryKeys()) == 0 {
		sql = c.insert(sc)
	} else {
		sql = c.upsert(sc)
	}
	vals := transformRecord(m.Record)
	for _, v := range vals {
		if _, err := c.db.ExecContext(ctx, sql, v...); err != nil {
			return fmt.Errorf("failed to execute '%s': %w", sql, err)
		}
	}
	return nil
}

func (*Client) insert(sc *arrow.Schema) string {
	var sb strings.Builder
	tableName, ok := sc.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		panic("missing table name in schema metadata")
	}
	sb.WriteString("insert into ")
	sb.WriteString(identifier(tableName))
	sb.WriteString(" (")
	columns := sc.Fields()
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(identifier(c.Name))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func (*Client) upsert(sc *arrow.Schema) string {
	var sb strings.Builder
	tableName, ok := sc.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		panic("missing table name in schema metadata")
	}
	sb.WriteString("insert or replace into ")
	sb.WriteString(identifier(tableName))
	sb.WriteString(" (")
	columns := sc.Fields()
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(identifier(c.Name))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}
