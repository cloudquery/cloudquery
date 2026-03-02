package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) Read(_ context.Context, table *schema.Table, res chan<- arrow.RecordBatch) error {
	sc := table.ToArrowSchema()
	index := c.Meilisearch.Index(table.Name)

	req := &meilisearch.SearchRequest{
		HitsPerPage: 100, // default = 1, we want more
		Page:        1,   // starting from 1
	}

	for {
		resp, err := index.Search("", req)
		if err != nil {
			return err
		}

		for _, hit := range resp.Hits {
			m := make(map[string]any)
			for k, v := range hit {
				var val any
				if err := json.Unmarshal(v, &val); err != nil {
					return fmt.Errorf("failed to unmarshal hit value for key %q: %w", k, err)
				}
				m[k] = val
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

func docToRecord(sc *arrow.Schema, doc map[string]any) (arrow.RecordBatch, error) {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)

	for i, builder := range builder.Fields() {
		if err := reverseTransform(builder, doc[sc.Field(i).Name]); err != nil {
			return nil, err
		}
	}

	return builder.NewRecordBatch(), nil
}
