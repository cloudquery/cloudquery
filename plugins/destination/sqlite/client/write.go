package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, res); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) (err error) {
	if len(msgs) == 0 {
		return nil
	}

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
			if err != nil {
				c.logger.Error().Err(err).Msg("failed to commit transaction")
			}
		}
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				c.logger.Error().Err(rollbackErr).Str("table", msgs[0].GetTable().Name).Msg("Failed to rollback transaction")
			}
		}
	}()

	for _, msg := range msgs {
		err = c.insertMessage(ctx, tx, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) insertMessage(ctx context.Context, tx *sql.Tx, m *message.WriteInsert) error {
	table := m.GetTable()
	sc := m.Record.Schema()
	var sqlString string
	if len(table.PrimaryKeys()) == 0 {
		sqlString = c.insert(sc)
	} else {
		sqlString = c.upsert(sc)
	}
	vals := transformRecord(m.Record)
	for _, v := range vals {
		if _, err := tx.ExecContext(ctx, sqlString, v...); err != nil {
			return fmt.Errorf("failed to execute '%s': %w", sqlString, err)
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
