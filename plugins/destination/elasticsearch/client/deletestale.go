package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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
func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	syncTimeStr := syncTime.Format(time.RFC3339)
	dateRange := types.NewDateRangeQuery()
	dateRange.Lt = &syncTimeStr
	q := types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{
				{
					MatchPhrase: map[string]types.MatchPhraseQuery{
						schema.CqSourceNameColumn.Name: {
							Query: source,
						},
					},
				},
				{
					Range: map[string]types.RangeQuery{
						schema.CqSyncTimeColumn.Name: dateRange,
					},
				},
			},
		},
	}
	req := deletebyquery.NewRequest()
	req.Query = &q

	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(maxConcurrentDeletes)
	for _, table := range tables {
		table := table
		g.Go(func() error {
			return c.deleteStaleTable(gctx, table, req, syncTime)
		})
	}
	return g.Wait()
}

func (c *Client) deleteStaleTable(ctx context.Context, table *schema.Table, req *deletebyquery.Request, syncTime time.Time) error {
	index := c.getIndexName(table.Name, syncTime)
	resp, err := c.typedClient.DeleteByQuery(index).Request(req).Do(ctx)
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
