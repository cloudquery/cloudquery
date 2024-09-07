package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"golang.org/x/exp/maps"
)

func getInsertQueryBuild(table *schema.Table, rowCount int, overwrite bool) *strings.Builder {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + identifier(table.Name))
	builder.WriteString(" (")

	for i, col := range table.Columns {
		builder.WriteString(identifier(col.Name))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES ")

	for i := 1; i <= rowCount; i++ {
		if i > 1 {
			builder.WriteString(", ")
		}
		builder.WriteString("(")
		builder.WriteString(strings.TrimSuffix(strings.Repeat("?,", len(table.Columns)), ","))
		builder.WriteString(")")
	}

	if overwrite {
		builder.WriteString(" ON DUPLICATE KEY UPDATE ")
		for i, col := range table.Columns {
			builder.WriteString(fmt.Sprintf("%s = VALUES(%s)", identifier(col.Name), identifier(col.Name)))
			if i < len(table.Columns)-1 {
				builder.WriteString(", ")
			}
		}
	}

	return &builder
}

func logTablesWithTruncation(logger zerolog.Logger, tables map[string]bool) {
	if len(tables) == 0 {
		return
	}
	keys := maps.Keys(tables)
	for k := range tables {
		keys = append(keys, k)
	}
	logger.Warn().Strs("tables", keys).Msg("tables contain a value in a primary key that is longer than what is supported by MySQL. only the first 191 characters will be included in the index. To see the complete record enable debug logs using `--log-level debug`")
}

func (c *Client) writeResources(ctx context.Context, overwrite bool, table *schema.Table, msgs message.WriteInserts) error {
	tablesWithTruncation := make(map[string]bool)
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

	rowCount := 0
	queryArgs := []any{}
	for _, msg := range msgs {
		rec := msg.Record
		transformedRecords, err := transformRecord(rec)
		if err != nil {
			return err
		}

		for _, record := range transformedRecords {
			for _, truncatablePKIndex := range pks {
				// log a warning that a blob or text field that is a PK has more than the limit
				lengthPerPk := c.maxIndexLength / len(pks)
				if len(record[truncatablePKIndex].(string)) > lengthPerPk {
					indexes := table.PrimaryKeysIndexes()
					pkValues := make(map[string]any, len(indexes))
					for i, pkIndex := range indexes {
						pkValues[table.Columns[pkIndex].Name] = record[i]
					}
					c.logger.Debug().Any("pk_values", pkValues).Msgf("record contains a primary key that is longer than MySQL can handle. only the first %d will be included in the index", lengthPerPk)
					tablesWithTruncation[table.Name] = true
					break
				}
			}
		}

		for _, transformedRecord := range transformedRecords {
			queryArgs = append(queryArgs, transformedRecord...)
			rowCount++
		}
	}

	query := getInsertQueryBuild(table, rowCount, overwrite)
	_, err := c.db.ExecContext(ctx, query.String(), queryArgs...)
	if err != nil {
		logTablesWithTruncation(c.logger, tablesWithTruncation)
		return err
	}
	logTablesWithTruncation(c.logger, tablesWithTruncation)
	return nil
}

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, resources message.WriteInserts) error {
	return c.writeResources(ctx, false, resources[0].GetTable(), resources)
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, msgs message.WriteInserts) error {
	return c.writeResources(ctx, true, msgs[0].GetTable(), msgs)
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
