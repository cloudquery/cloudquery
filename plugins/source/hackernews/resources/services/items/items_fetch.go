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

// fetchItems gets called by the CloudQuery SDK to fetch the data for the `hackernews_items` table. It is passed a context and a
// schema.ClientMeta interface which is used to access the plugin configuration and the state backend. The
// resource argument is not used here, as it refers to the parent table, which is not applicable in this case.
// The res channel is used to send the fetched data back to the CloudQuery SDK.
//
// Because this is an incremental table, we must first load the cursor from the state backend, and then update
// it whenever we've successfully fetched a new item. The cursor is a string, and in this case it is the ID of
// the last item we've fetched.
//
// Incremental tables should guarantee at-least-once delivery. The Hacker News API gives us the current maximum item ID,
// and we need to fetch all items up to that ID. Once we've fetched all items up to a certain ID, we can safely
// update the cursor. If we crash before updating the cursor, we will fetch the same items again on the next run.
//
// This resolver is different from most in that it needs to concurrently fetch items using its own goroutines. Usually,
// a resolver with pagination will not have this ability and will simply iterate over all pages.
func fetchItems(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	tableName := Items().Name
	value, err := c.Backend.Get(ctx, tableName, c.ID())
	if err != nil {
		return fmt.Errorf("failed to retrieve state from backend: %w", err)
	}

	// read the cursor from the state, or default to 0 if it's not set
	cursor := 0
	if value == "" {
		c.Logger().Info().
			Str("table", tableName).
			Str("client_id", c.ID()).
			Msgf("No previous cursor found")
	} else {
		cursor, err = strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("failed to convert cursor to int: %w", err)
		}
		c.Logger().Info().
			Str("table", tableName).
			Str("client_id", c.ID()).
			Msg("Found previous cursor with value " + strconv.Itoa(cursor))
	}

	// find the max item ID from the Hacker News API
	work := make(chan int)
	maxID, err := c.HackerNews.MaxItemID(ctx)
	if err != nil {
		return err
	}
	c.Logger().Info().Msg("Found max ID, reading up to " + strconv.Itoa(maxID))

	// spin up worker goroutines that will fetch items concurrently by reading ids from
	// the work channel
	workers := c.Spec.ItemConcurrency
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

	// We use a min heap to keep track of the items that have finished fetching,
	// but are ahead of where the cursor is. This allows us to guarantee at-least-once
	// delivery of items, because the cursor is only moved on once we are sure all items
	// up to the new position have been synced, and there are no gaps.
	//
	// This diagram attempts to explain the algorithm visually:
	//
	// itemID      | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
	// success     | y | y | ? | ? | y | y | y |   |   |
	// cursor      |   | x |   |   |   |   |   |   |   |
	// heap        |   |   |   |   | x | x | x |   |   |
	//
	// When item 3 is done, it will get added to the heap, but then immediately removed when the cursor
	// is moved to position 3. When item 4 is done, it will also get added to the heap, but the cursor
	// will now get moved to position 7, and the heap will be emptied.
	g2, g2ctx := errgroup.WithContext(ctx)
	g2.Go(func() error {
		h := &intheap.IntHeap{}
		for v := range success {
			heap.Push(h, v)
			for h.Len() > 0 {
				min := heap.Pop(h).(int)
				if min != cursor+1 {
					heap.Push(h, min)
					break
				}
				cursor++
			}
			err = c.Backend.Set(g2ctx, tableName, c.ID(), strconv.Itoa(cursor))
			if err != nil {
				return fmt.Errorf("failed to update state backend: %w", err)
			}
		}
		return nil
	})

	// send work to the workers until we reach maxID.
	for i := cursor + 1; i <= maxID; i++ {
		select {
		case work <- i:
		case <-gctx.Done():
			close(work)
			return gctx.Err()
		}
	}
	close(work)

	// after the work channel is closed, we wait for the workers to finish
	err = g.Wait()
	if err != nil {
		return err
	}

	// now we can close the success channel, and wait for the final cursor value
	// to be set
	close(success)
	return g2.Wait()
}

// fetchItem fetches a single item from the Hacker News API and sends it to the CloudQuery SDK
// to dispatch to destinations via the res channel.
func fetchItem(ctx context.Context, c *client.Client, itemID int, res chan<- any) error {
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
