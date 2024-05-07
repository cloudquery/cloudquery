package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

func (c Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	filtered, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	for _, table := range filtered {
		res <- &message.SyncMigrateTable{
			Table: table,
		}
	}
	return c.syncTables(ctx, filtered, res)
}

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- message.SyncMessage) error {
	var err error
	lastDocumentId := ""
	maxBatchSize := c.maxBatchSize
	collection := c.client.Collection(table.Name)
	for {
		orderBy := firestore.DocumentID
		if c.orderBy != "" {
			orderBy = c.orderBy
		}

		dir := firestore.Asc
		if c.orderDirection == "desc" {
			dir = firestore.Desc
		}

		query := collection.Query.
			OrderBy(orderBy, dir).
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

			arrowSchema := table.ToArrowSchema()
			rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
			idField := rb.Field(0).(*array.StringBuilder)
			idField.AppendString(docSnap.Ref.ID)

			createdAtField := rb.Field(1).(*array.TimestampBuilder)
			createdAtField.AppendTime(docSnap.CreateTime)

			updatedAtField := rb.Field(2).(*array.TimestampBuilder)
			updatedAtField.AppendTime(docSnap.UpdateTime)

			dataField := rb.Field(3).(*types.JSONBuilder)
			dataField.Append(docSnap.Data())

			res <- &message.SyncInsert{Record: rb.NewRecord()}
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

func (c *Client) syncTables(ctx context.Context, tables schema.Tables, res chan<- message.SyncMessage) error {
	eg, gctx := errgroup.WithContext(ctx)
	for _, table := range tables {
		t := table
		eg.Go(func() error {
			return c.syncTable(gctx, t, res)
		})
	}
	return eg.Wait()
}
