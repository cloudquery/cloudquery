package regions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Regions() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_regions",
		Resolver:  fetchRegions,
		Transform: transformers.TransformWithStruct(&godo.Region{}, transformers.WithPrimaryKeys("Slug")),
	}
}

func fetchRegions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		var (
			items []godo.Region
			resp  *godo.Response
		)
		getFunc := func() error {
			var err error
			items, resp, err = svc.Services.Regions.List(ctx, opt)
			return err
		}
		err := client.ThrottleWrapper(ctx, svc, getFunc)
		if err != nil {
			return err
		}
		res <- items
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		opt.Page = page + 1
	}
	return nil
}
