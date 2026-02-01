package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/deletebyquery"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"golang.org/x/sync/errgroup"
)

type deleteByQueryErrorResponse struct {
	Status int `json:"status"`
	Error  struct {
		RootCause []struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"root_cause"`
	} `json:"error"`
}

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
	resp, err := c.typedClient.DeleteByQuery(index).Request(req).WaitForCompletion(true).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete stale entries: %w", err)
	}
	defer resp.Body.Close()
	var qResp deleteByQueryErrorResponse
	b, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(b, &qResp); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if qResp.Status != 0 {
		if len(qResp.Error.RootCause) > 0 {
			if qResp.Error.RootCause[0].Type == "index_not_found_exception" {
				return nil
			}
			return fmt.Errorf("failed to delete stale entries: %s", qResp.Error.RootCause[0].Reason)
		}
		return fmt.Errorf("failed to delete stale entries: status %d", qResp.Status)
	}
	return nil
}
