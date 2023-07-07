package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

func getInsertQueryBuild(table *schema.Table) *strings.Builder {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + identifier(table.Name))
	builder.WriteString(" (")

	for i, col := range table.Columns {
		builder.WriteString(identifier(col.Name))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.TrimSuffix(strings.Repeat("?,", len(table.Columns)), ","))
	builder.WriteString(")")
	return &builder
}

func logTablesWithTruncation(logger zerolog.Logger, tables map[string]bool) {
	keys := make([]string, len(tables))
	if len(tables) == 0 {
		return
	}
	for k := range tables {
		keys = append(keys, k)
	}
	logger.Warn().Strs("tables", keys).Msg("tables contain a value in a primary key that is longer than what is supported by MySQL. only the first 191 characters will be included in the index. To see the complete record enable debug logs using `--log-level debug`")
}

func (c *Client) writeResources(ctx context.Context, query string, msgs message.WriteInserts) error {
	tablesWithTruncation := make(map[string]bool)
	for _, msg := range msgs {
		// This assumes that all the records in the message are from the same table
		// If we can assume all msgs are from the same table, we can move this outside the loop
		table := msg.GetTable()
		pks := make([]int, 0)
		for i, col := range table.Columns {
			if !col.PrimaryKey {
				continue
			}
			sqlType := arrowTypeToMySqlStr(col.Type)
			if sqlType != "blob" && sqlType != "text" {
				continue
			}
			// only if the PK is a blob or a text do we care about the length of the data
			pks = append(pks, i)
		}
		rec := msg.Record
		transformedRecords, err := transformRecord(rec)
		if err != nil {
			return err
		}

		// log a warning that a blob or text field that is a PK has more than 191 characters
		for _, record := range transformedRecords {
			for _, pki := range pks {
				if len(record[pki].(string)) > maxPrefixLength {
					c.logger.Debug().Any("record", record).Msgf("record contains a primary key that is longer than MySQL can handle. only the first %d will be included in the index", maxPrefixLength)
					tablesWithTruncation[table.Name] = true
				}
			}
		}

		for _, transformedRecord := range transformedRecords {
			_, err := c.db.ExecContext(ctx, query, transformedRecord...)
			if err != nil {
				logTablesWithTruncation(c.logger, tablesWithTruncation)
				return err
			}
		}
	}
	logTablesWithTruncation(c.logger, tablesWithTruncation)
	return nil
}

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, resources message.WriteInserts) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(";")
	return c.writeResources(ctx, builder.String(), resources)
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, msgs message.WriteInserts) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, col := range table.Columns {
		builder.WriteString(fmt.Sprintf("%s = VALUES(%s)", identifier(col.Name), identifier(col.Name)))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	return c.writeResources(ctx, builder.String(), msgs)
}

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, res); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	if len(msgs) == 0 {
		return nil
	}
	table := msgs[0].GetTable()
	hasPks := len(table.PrimaryKeys()) > 0
	if hasPks {
		return c.overwriteTableBatch(ctx, table, msgs)
	}
	return c.appendTableBatch(ctx, table, msgs)
}
