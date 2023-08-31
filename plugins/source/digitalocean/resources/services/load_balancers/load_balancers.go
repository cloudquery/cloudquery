package load_balancers

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_load_balancers",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Load-Balancers",
		Resolver:    fetchLoadBalancers,
		Transform:   transformers.TransformWithStruct(&godo.LoadBalancer{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		var (
			data []godo.LoadBalancer
			resp *godo.Response
		)
		listFunc := func() error {
			var err error
			data, resp, err = svc.Services.LoadBalancers.List(ctx, opt)
			return err
		}
		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return err
		}
		res <- data
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
