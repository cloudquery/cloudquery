package load_balancers

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_load_balancers",
		Resolver:  fetchLoadBalancers,
		Transform: transformers.TransformWithStruct(&godo.LoadBalancer{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		var (
			items []godo.LoadBalancer
			resp  *godo.Response
		)
		getFunc := func() error {
			var err error
			items, resp, err = svc.Services.LoadBalancers.List(ctx, opt)
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
