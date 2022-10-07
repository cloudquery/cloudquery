package hooks

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func fetchDeliveries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	id := *parent.Item.(*github.Hook).ID

	c := meta.(*client.Client)
	opts := &github.ListCursorOptions{PerPage: 100}

	for {
		deliveries, resp, err := c.Github.Organizations.ListHookDeliveries(ctx, c.Org, id, opts)
		if err != nil {
			return err
		}
		res <- deliveries

		opts.Cursor = resp.NextPageToken
		if resp.NextPageToken == "" {
			return nil
		}
	}
}

func resolveRequest(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	delivery := resource.Item.(*github.HookDelivery)
	return resource.Set(c.Name, delivery.Request.String())
}

func resolveResponse(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	delivery := resource.Item.(*github.HookDelivery)
	return resource.Set(c.Name, delivery.Response.String())
}
