package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	index := c.getIndexNamePattern(table.Name)

	// refresh index before read, to ensure all written data is available
	resp, err := c.typedClient.Indices.Refresh().Index(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to refresh index before read: %w", err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// do the read
	resp, err = c.typedClient.Search().Index(index).Request(&search.Request{
		Query: &types.Query{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				schema.CqSourceNameColumn.Name: {
					Query: sourceName,
				},
			},
		},
	}).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Hits struct {
			Hits []struct {
				Source map[string]any `json:"_source"`
			}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	for _, hit := range result.Hits.Hits {
		values := make([]any, len(table.Columns))
		for i, col := range table.Columns {
			switch col.Type {
			case schema.TypeByteArray:
				ba, err := base64.StdEncoding.DecodeString(hit.Source[col.Name].(string))
				if err != nil {
					return fmt.Errorf("failed to decode base64 string: %w", err)
				}
				values[i] = ba
			default:
				values[i] = hit.Source[col.Name]
			}
		}
		res <- values
	}
	return nil
}
