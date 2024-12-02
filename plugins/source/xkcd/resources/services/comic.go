package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/xkcd/client"
	"github.com/cloudquery/cloudquery/plugins/source/xkcd/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/sync/errgroup"
)

func ComicsTable() *schema.Table {
	return &schema.Table{
		Name:      "xkcd_comics",
		Resolver:  fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}, transformers.WithPrimaryKeys("Num")),
	}
}

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	comic, err := c.XKCD.GetLatestComic(0)
	if err != nil {
		return err
	}
	res <- comic
	g := errgroup.Group{}
	g.SetLimit(10)
	for i := 1; i < comic.Num; i++ {
		if i == 404 {
			continue
		}
		i := i
		g.Go(func() error {
			comic, err := c.XKCD.GetComic(i)
			if err != nil {
				c.Logger.Error().Err(err).Msgf("failed to fetch comic %d", i)
				return err
			}
			res <- comic
			return nil
		})
	}
	return g.Wait()
}
