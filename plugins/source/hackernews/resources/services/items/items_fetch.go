package items

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/internal/intheap"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/hermanschaaf/hackernews"
	"golang.org/x/sync/errgroup"
)

func fetchItems(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	tableName := Items().Name
	value, err := c.Backend.Get(tableName, "id")
	if err != nil {
		return fmt.Errorf("failed to retrieve state from backend: %w", err)
	}
	cursor := 0
	if value != "" {
		cursor, err = strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("failed to convert cursor to int: %w", err)
		}
	}

	c.Logger().Info().Msg("Found previous cursor, starting from " + strconv.Itoa(cursor))

	work := make(chan int)
	maxID, err := c.HackerNews.MaxItemID(ctx)
	if err != nil {
		return err
	}

	c.Logger().Info().Msg("Found max ID, reading up to " + strconv.Itoa(maxID))

	workers := c.Spec.ItemConcurrency
	if workers == 0 {
		workers = 5
	}

	success := make(chan int)
	g, gctx := errgroup.WithContext(ctx)
	for i := 0; i < workers; i++ {
		g.Go(func() error {
			for {
				select {
				case <-gctx.Done():
					return gctx.Err()
				case itemID, ok := <-work:
					if !ok {
						return nil
					}
					fetchErr := c.RetryOnError(ctx, tableName, func() error {
						return fetchItem(gctx, c, itemID, res)
					})
					if fetchErr != nil {
						return fetchErr
					}
					success <- itemID
				}
			}
		})
	}

	// We use a min heap to keep track of the smallest item id that we have not yet fetched.
	// The cursor should always be one less than this item id.

	g2 := &errgroup.Group{}
	g2.Go(func() error {
		h := &intheap.IntHeap{}
		for v := range success {
			c.Logger().Info().Msg("Done with" + strconv.Itoa(v))
			heap.Push(h, v)
			for h.Len() > 0 {
				min := heap.Pop(h).(int)
				if min != cursor+1 {
					heap.Push(h, min)
					break
				}
				cursor++
			}
			c.Backend.Set(tableName, "id", strconv.Itoa(cursor))
		}
		return nil
	})
	for i := cursor + 1; i <= maxID; i++ {
		c.Logger().Info().Msg("Queuing item " + strconv.Itoa(i))
		select {
		case work <- i:
		case <-gctx.Done():
			close(work)
			return gctx.Err()
		}
	}
	close(work)
	err = g.Wait()
	if err != nil {
		return err
	}
	close(success)
	return g2.Wait()
}

func fetchItem(ctx context.Context, c *client.Client, itemID int, res chan<- any) error {
	c.Logger().Info().Msg("Fetching item " + strconv.Itoa(itemID))
	item, err := c.HackerNews.GetItem(ctx, itemID)
	if err != nil {
		var httpErr hackernews.HTTPError
		if errors.As(err, &httpErr) && httpErr.Code == 404 {
			// Assume item was deleted
			return nil
		}
		return err
	}
	res <- item
	return nil
}
