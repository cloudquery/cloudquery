package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/sync/errgroup"
)

const maxConcurrentDeletes = 10

// DeleteStale removes entries from previous syncs
func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(maxConcurrentDeletes)
	for _, msg := range msgs {
		msg := msg
		g.Go(func() error {
			syncTimeStr := msg.SyncTime.Format(time.RFC3339)
			return c.deleteStaleIndex(gctx,
				msg.TableName,
				&deletebyquery.Request{
					Query: &types.Query{
						Bool: &types.BoolQuery{
							Filter: []types.Query{
								{MatchPhrase: map[string]types.MatchPhraseQuery{schema.CqSourceNameColumn.Name: {Query: msg.SourceName}}},
								{Range: map[string]types.RangeQuery{schema.CqSyncTimeColumn.Name: &types.DateRangeQuery{Lt: &syncTimeStr}}},
							},
						},
					},
				},
			)
		})
	}
	return g.Wait()
}

func (c *Client) deleteStaleIndex(ctx context.Context, index string, req *deletebyquery.Request) error {
	_, err := c.typedClient.DeleteByQuery(index).Request(req).WaitForCompletion(true).Do(ctx)
	if err != nil {
		var esErr *types.ElasticsearchError
		if errors.As(err, &esErr) && esErr.ErrorCause.Type == "index_not_found_exception" {
			return nil
		}
		return fmt.Errorf("failed to delete stale entries: %w", err)
	}
	return nil
}
