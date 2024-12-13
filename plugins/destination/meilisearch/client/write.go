package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, res); err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	return c.writer.Flush(ctx)
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, messages message.WriteInserts) error {
	if len(messages) == 0 {
		return nil
	}
	table := messages[0].GetTable()

	index, err := c.Meilisearch.GetIndex(name)
	if err != nil {
		return err
	}

	transformer := transform(table)

	records := make([]arrow.Record, 0, len(messages))
	for _, msg := range messages {
		records = append(records, msg.Record)
	}

	docs := make([]map[string]any, 0, len(records)) // at least 1 row in record
	for _, record := range records {
		rows, err := transformer(record)
		if err != nil {
			return err
		}
		docs = append(docs, rows...)
	}

	taskInfo, err := index.AddDocuments(&docs, index.PrimaryKey)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to write %d items to index %q: %w", len(docs), index.UID, err)
	}

	return nil
}

type rowTransformer func(record arrow.Record) ([]map[string]any, error)

func toMap(table *schema.Table) rowTransformer {
	columns := table.Columns.Names()
	return func(record arrow.Record) ([]map[string]any, error) {
		byColumn := make(map[string][]any, len(columns))
		for i, col := range record.Columns() {
			byColumn[columns[i]] = getValues(col)
		}
		return transpose(byColumn, int(record.NumRows())), nil
	}
}

func transform(table *schema.Table) rowTransformer {
	m := toMap(table)
	h := hashUUID(table)
	// we always use the hashUUID func
	return func(record arrow.Record) ([]map[string]any, error) {
		rows, err := m(record)
		if err != nil {
			return nil, err
		}
		for _, row := range rows {
			row[hashColumnName] = h(row)
		}
		return rows, nil
	}
}
