package client

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/json"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// func (c *Client) Write(ctx context.Context, tables schema.Schemas, res <-chan arrow.Record) error {
// 	var sql string
// 	for r := range res {

// 		tableName := schema.TableName(r.Schema())
// 		table := tables.SchemaByName(tableName)
// 		if c.spec.WriteMode == specs.WriteModeAppend {
// 			sql = c.insertSQL(table, r.Data)
// 			expanded := expandData(table, r.Data)
// 			connector, _ := duckdb.NewConnector("", nil)
// 			conn, _ := connector.Connect(ctx)
// 			appender, _ := duckdb.NewAppenderFromConn(conn, "", "")
// 			appender.
// 			appender.AppendRow()
// 			if _, err := c.db.Exec(sql, expanded...); err != nil {
// 				return fmt.Errorf("failed to execute '%s': %w", sql, err)
// 			}
// 		} else {
// 			err := c.upsert(table, r.Data)
// 			if err != nil {
// 				return fmt.Errorf("failed to execute upsert: %w", err)
// 			}
// 		}
// 	}
// 	return nil
// }

func expandData(table *arrow.Schema, data []any) []any {
	var expanded []any
	for i, d := range data {
		if arrow.IsListLike(table.Field(i).Type.ID()) {
			arr := d.([]any)
			for _, v := range arr {
				if v != nil {
					expanded = append(expanded, v)
				}
			}
			// expanded = append(expanded, arr...)
		} else {
			expanded = append(expanded, d)
		}
	}
	return expanded
}

// func (c *Client) insertSQL(record arrow.Record) string {
// 	var sb strings.Builder
// 	sb.WriteString("insert into ")
// 	c.insertQuery(&sb, table, data)
// 	return sb.String()
// }

// func (c *Client) upsert(table *schema.Table, data []any) error {
// 	// At time of writing (March 2023), duckdb does not support updating list columns.
// 	// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
// 	// but this is unavoidable until support is added to duckdb itself.
// 	// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
// 	var sb strings.Builder
// 	if len(table.PrimaryKeys()) > 0 {
// 		sb.WriteString("delete from ")
// 		sb.WriteString(`"` + table.Name + `"`)
// 		sb.WriteString(" where ")
// 		pks := table.PrimaryKeys()
// 		pkData := make([]any, len(pks))
// 		for i, k := range pks {
// 			col := table.Columns.Get(k)
// 			sb.WriteString(`"` + col.Name + `"`)
// 			sb.WriteString(" = ")
// 			sb.WriteString(fmt.Sprintf("$%d ", i+1))
// 			if i < len(pks)-1 {
// 				sb.WriteString("and ")
// 			}
// 			pkData[i] = data[table.Columns.Index(k)]
// 		}
// 		sql := sb.String()
// 		if _, err := c.db.Exec(sql, pkData...); err != nil {
// 			return fmt.Errorf("failed to execute '%s': %w", sql, err)
// 		}
// 		sb.Reset()
// 	}
// 	sb.WriteString("insert into ")
// 	c.insertQuery(&sb, table, data)
// 	expanded := expandData(table, data)
// 	sql := sb.String()
// 	if _, err := c.db.Exec(sql, expanded...); err != nil {
// 		return fmt.Errorf("failed to execute '%s': %w", sql, err)
// 	}
// 	return nil
// }

func (*Client) insertQuery(table *arrow.Schema, data []any) string {
	var sb strings.Builder
	tableName := schema.TableName(table)
	sb.WriteString("insert into ")
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" (")
	columns := table.Fields()
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	counter := 0
	for i := range columns {
		if arrow.IsListLike(table.Field(i).Type.ID()) {
			n := len(data[i].([]any))
			sb.WriteString("[")
			for j := 0; j < n; j++ {
				if data[i].([]any)[j] == nil {
					sb.WriteString("null")
				} else {
					sb.WriteString(fmt.Sprintf("$%d", counter+1))
					counter++
				}
				// sb.WriteString(fmt.Sprintf("$%d", counter+1))
				// counter++
				if j < n-1 {
					sb.WriteString(",")
				}
			}
			sb.WriteString("]")
		} else {
			sb.WriteString(fmt.Sprintf("$%d", counter+1))
			counter++
		}
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}


func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	tableName := schema.TableName(sc)
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.json", tableName))
	if err != nil {
		return err
	}
	// defer os.Remove(f.Name())

	w := json.NewWriter(f, sc)
	for _, r := range records {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	
	if err := f.Close(); err != nil {
		return err
	}
	_, err = c.db.Exec("copy " + tableName + " from '" + f.Name() + "' (timestampformat '%Y-%m-%d %H:%M:%S.%f')")
	if err != nil {
		return err
	}
	return nil
}