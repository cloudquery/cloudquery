package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	index, err := c.Meilisearch.GetIndex(table.Name)
	if err != nil {
		return err
	}

	var transformer rowTransformer
	switch c.dstSpec.WriteMode {
	case specs.WriteModeAppend:
		transformer = toMap(table)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		transformer = toMapWithHash(table)
	default:
		return fmt.Errorf("unsupported write mode %q", c.dstSpec.WriteMode.String())
	}

	docs := make([]map[string]any, len(data))
	for i, item := range data {
		docs[i] = transformer(item)
	}

	taskInfo, err := index.AddDocuments(&docs, c.pkColumn)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to write %d items to index %q: %w", len(data), index.UID, err)
	}

	return nil
}

type rowTransformer func(item []any) map[string]any

func toMap(table *schema.Table) rowTransformer {
	columns := table.Columns.Names()
	return func(item []any) map[string]any {
		res := make(map[string]any, len(columns))
		for i, col := range columns {
			res[col] = item[i]
		}
		return res
	}
}

func toMapWithHash(table *schema.Table) rowTransformer {
	m := toMap(table)
	h := hashUUID(table)
	return func(item []any) map[string]any {
		res := m(item)
		res[hashColumnName] = h(item)
		return res
	}
}
