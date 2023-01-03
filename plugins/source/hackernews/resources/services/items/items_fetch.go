package items

import (
	"context"
	"errors"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/hermanschaaf/hackernews"
	"golang.org/x/sync/errgroup"
)

func fetchItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	tableName := Items().Name
	cursor := c.State.Get(tableName, "id")
	var work chan int

	maxID, err := c.HackerNews.MaxItemID(ctx)
	if err != nil {
		return err
	}

	g, gtcx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			select {
			case <-gtcx.Done():
				return gtcx.Err()
			case itemID := <-work:
				f := func() error {
					return fetchItem(gtcx, c, itemID, res)
				}
				return c.RetryOnError(ctx, tableName, f)
			}
		}
	})

	for i := cursor + 1; i <= maxID; i++ {
		work <- i
	}
}

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
