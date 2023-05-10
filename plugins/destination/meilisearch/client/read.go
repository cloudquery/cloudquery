package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) Read(_ context.Context, sc *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	index, err := c.Meilisearch.GetIndex(schema.TableName(sc))
	if err != nil {
		return err
	}

	req := &meilisearch.SearchRequest{
		Filter:      schema.CqSourceNameColumn.Name + " = '" + sourceName + "'",
		Sort:        []string{schema.CqSyncTimeColumn.Name + ":asc"},
		HitsPerPage: 100, // default = 1, we want more
		Page:        1,   // starting from 1
	}

	for {
		resp, err := index.Search("", req)
		if err != nil {
			return err
		}

		for _, hit := range resp.Hits {
			m, ok := hit.(map[string]any)
			if !ok {
				return fmt.Errorf("unsupported format for doc: %T", hit)
			}
			row, err := docToRecord(sc, m)
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

func docToRecord(sc *arrow.Schema, doc map[string]any) (arrow.Record, error) {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)

	for i, builder := range builder.Fields() {
		if err := reverseTransform(builder, doc[sc.Field(i).Name]); err != nil {
			return nil, err
		}
	}

	return builder.NewRecord(), nil
}
