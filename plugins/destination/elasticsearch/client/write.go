package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/segmentio/fasthash/fnv1a"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) (err error) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         fmt.Sprintf("%s-%s", table.Name, time.Now().Format("2006-01-02")), // The default index name
		Client:        c.client,                                                          // The Elasticsearch client
		NumWorkers:    c.pluginSpec.Concurrency,                                          // The number of worker goroutines
		FlushBytes:    c.spec.BatchSizeBytes,                                             // The flush threshold in bytes
		FlushInterval: 30 * time.Second,                                                  // The periodic flush interval
	})
	if err != nil {
		return err
	}
	defer func() {
		err = bi.Close(ctx)
	}()
	failed := false
	for _, r := range resources {
		if failed {
			break
		}
		doc := map[string]any{}
		for i, col := range table.Columns {
			doc[col.Name] = r[i]
		}
		data, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		docID := resourceID(table, r)
		c.logger.Info().Msgf("Writing resource %s/%d", table.Name, docID)
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: fmt.Sprintf("%d", docID),
				Body:       bytes.NewReader(data),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					failed = true
					l := c.logger.Error().Err(err)
					if err != nil {
						l.Err(err)
					} else {
						c.logger.Error().Str("error_type", res.Error.Type).Str("error_reason", res.Error.Reason)
					}
					l.Msg("Unexpected error while indexing document")
				},
			},
		)
		if err != nil {
			return fmt.Errorf("failed to add item to bulk indexer: %w", err)
		}
	}
	return
}

// elasticsearch IDs are limited to 512 bytes, so we hash the resource PK to make sure it's within the limit
func resourceID(table *schema.Table, resource []any) uint64 {
	parts := make([]string, 0, len(table.PrimaryKeys()))
	for i, col := range table.Columns {
		if !col.CreationOptions.PrimaryKey {
			continue
		}
		parts = append(parts, fmt.Sprint(resource[i]))
	}
	h1 := fnv1a.HashString64(strings.Join(parts, "-"))
	return h1
}
