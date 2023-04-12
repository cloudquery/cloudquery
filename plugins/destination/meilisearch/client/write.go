package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func releaseRecords(records []arrow.Record) {
	for _, record := range records {
		record.Release()
	}
}

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	defer releaseRecords(records)

	index, err := c.Meilisearch.GetIndex(schema.TableName(sc))
	if err != nil {
		return err
	}

	var transformer rowTransformer
	switch c.dstSpec.WriteMode {
	case specs.WriteModeAppend:
		transformer = toMap(sc)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		transformer = toMapWithHash(sc)
	default:
		return fmt.Errorf("unsupported write mode %q", c.dstSpec.WriteMode.String())
	}

	docs := make([]map[string]any, 0, len(records)) // at least 1 row in record
	for _, record := range records {
		rows, err := transformer(record)
		if err != nil {
			return err
		}
		docs = append(docs, rows...)
	}

	taskInfo, err := index.AddDocuments(&docs, c.pkColumn)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to write %d items to index %q: %w", len(docs), index.UID, err)
	}

	return nil
}

type rowTransformer func(record arrow.Record) ([]map[string]any, error)

func toMap(sc *arrow.Schema) rowTransformer {
	columns := make([]string, len(sc.Fields()))
	for i, fld := range sc.Fields() {
		columns[i] = fld.Name
	}
	return func(record arrow.Record) ([]map[string]any, error) {
		byColumn := make(map[string][]any, len(columns))
		for i, col := range record.Columns() {
			byColumn[columns[i]] = getValues(col)
		}
		return transpose(byColumn, int(record.NumRows())), nil
	}
}

func toMapWithHash(sc *arrow.Schema) rowTransformer {
	m := toMap(sc)
	h := hashUUID(sc)
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
