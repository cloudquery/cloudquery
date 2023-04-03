package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) Read(_ context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	index, err := c.Meilisearch.GetIndex(table.Name)
	if err != nil {
		return err
	}

	req := &meilisearch.SearchRequest{
		Filter:      schema.CqSourceNameColumn.Name + " = '" + sourceName + "'",
		Sort:        []string{schema.CqSyncTimeColumn.Name + ":asc"},
		HitsPerPage: 100, // default = 1, we want more
		Page:        1,   // starting from 1
	}

	unwrap := unmap(table)
	for {
		resp, err := index.Search("", req)
		if err != nil {
			return err
		}

		for _, hit := range resp.Hits {
			row, err := unwrap(hit)
			if err != nil {
				return err
			}
			res <- row
		}

		if resp.TotalPages == req.Page {
			break
		}
		req.Page++
	}

	return nil
}

func unmap(table *schema.Table) func(hit any) ([]any, error) {
	return func(hit any) ([]any, error) {
		m, ok := hit.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unsupported format for doc: %T", hit)
		}

		res := make([]any, len(table.Columns))
		for idx, col := range table.Columns {
			res[idx] = m[col.Name]
		}
		return res, nil
	}
}
