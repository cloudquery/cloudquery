package client

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

func (c *Client) Sync(ctx context.Context, metrics *source.Metrics, res chan<- *schema.Resource) error {
	c.metrics = metrics
	for _, table := range c.Tables {
		if c.metrics.TableClient[table.Name] == nil {
			c.metrics.TableClient[table.Name] = make(map[string]*source.TableClientMetrics)
			c.metrics.TableClient[table.Name][c.ID()] = &source.TableClientMetrics{}
		}
	}
	return c.syncTables(ctx, res)
}

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- *schema.Resource) error {
	var err error
	lastDocumentId := ""
	maxBatchSize := c.maxBatchSize
	collection := c.client.Collection(table.Name)
	for {
		orderByField := firestore.DocumentID
		if c.orderByField != "" {
			orderByField = c.orderByField
		}

		dir := firestore.Asc
		if c.orderByDirection == "desc" {
			dir = firestore.Desc
		}

		query := collection.Query.
			OrderBy(orderByField, dir).
			Limit(maxBatchSize)
		if lastDocumentId != "" {
			c.logger.Info().Msgf("Starting after %s", lastDocumentId)
			query = query.StartAfter(lastDocumentId)
		}
		docIter := query.Documents(ctx)
		var documentCount int
		var skippedCount int
		for {
			docSnap, err := docIter.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}
				return err
			}
			documentCount++
			if !docSnap.Exists() {
				skippedCount++
				continue
			}
			lastDocumentId = docSnap.Ref.ID
			item := docSnap.Data()
			resource := schema.NewResourceData(table, nil, item)
			err = resource.Set("__id", docSnap.Ref.ID)
			if err != nil {
				return err
			}
			err = resource.Set("__created_at", docSnap.CreateTime)
			if err != nil {
				return err
			}
			err = resource.Set("__updated_at", docSnap.UpdateTime)
			if err != nil {
				return err
			}
			err = resource.Set("data", item)
			if err != nil {
				return err
			}
			c.metrics.TableClient[table.Name][c.ID()].Resources++
			res <- resource
		}
		c.logger.Info().Msgf("Synced %d documents from %s", documentCount, table.Name)
		if skippedCount > 0 {
			c.logger.Info().Msgf("Skipped %d documents from %s", skippedCount, table.Name)
		}
		if documentCount < maxBatchSize {
			break
		}
	}
	return err
}

func (c *Client) syncTables(ctx context.Context, res chan<- *schema.Resource) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(len(c.Tables))

	for _, table := range c.Tables {
		t := table
		eg.Go(func() error {
			return c.syncTable(ctx, t, res)
		})
	}
	return eg.Wait()
}
