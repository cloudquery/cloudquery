package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/segmentio/fasthash/fnv1a"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	if len(resources) == 0 {
		return nil
	}

	var buf bytes.Buffer
	pks := pkIndexes(table) // do some work up front to avoid doing it for every resource
	// get the sync time from the first resource in the batch (here we assume that all resources in the batch
	// have the same sync time. At the moment this assumption holds.)
	syncTime := resources[0][table.Columns.Index(schema.CqSyncTimeColumn.Name)].(time.Time)
	for _, r := range resources {
		doc := map[string]any{}
		for i, col := range table.Columns {
			doc[col.Name] = r[i]
		}
		data, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		var meta []byte
		if c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale {
			docID := fmt.Sprint(resourceID(r, pks))
			meta = []byte(fmt.Sprintf(`{"index":{"_id":"%s"}}%s`, docID, "\n"))
		} else {
			meta = []byte(`{"index":{}}` + "\n")
		}
		data = append(data, "\n"...)
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}
	index := c.getIndexName(table.Name, syncTime)
	resp, err := c.client.Bulk(bytes.NewReader(buf.Bytes()),
		c.client.Bulk.WithContext(ctx),
		c.client.Bulk.WithIndex(index),
	)
	if err != nil {
		return fmt.Errorf("failed to create bulk request: %w", err)
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return fmt.Errorf("bulk request failed: %s", resp.String())
	}
	return err
}

func pkIndexes(table *schema.Table) []int {
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// if no PK is defined, use all columns for the ID which is based on the indices returned by this function
		pks = table.Columns.Names()
	}
	inds := make([]int, 0, len(pks))
	for _, col := range pks {
		inds = append(inds, table.Columns.Index(col))
	}
	return inds
}

// elasticsearch IDs are limited to 512 bytes, so we hash the resource PK to make sure it's within the limit
func resourceID(resource []any, inds []int) uint64 {
	parts := make([]string, 0, len(inds))
	for _, i := range inds {
		parts = append(parts, fmt.Sprint(resource[i]))
	}
	h1 := fnv1a.HashString64(strings.Join(parts, "-"))
	return h1
}
