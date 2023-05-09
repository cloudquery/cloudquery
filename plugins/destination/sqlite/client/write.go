package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Write(ctx context.Context, tables schema.Schemas, res <-chan arrow.Record) error {
	var sql string
	for r := range res {
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insert(r.Schema())
		} else {
			sql = c.upsert(r.Schema())
		}
		vals := transformRecord(r)
		for _, v := range vals {
			if _, err := c.db.Exec(sql, v...); err != nil {
				return fmt.Errorf("failed to execute '%s': %w", sql, err)
			}
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
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" (")
	columns := sc.Fields()
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
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
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" (")
	columns := sc.Fields()
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
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
