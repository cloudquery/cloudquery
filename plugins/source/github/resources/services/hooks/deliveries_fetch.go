package hooks

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchDeliveries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

func hooksGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	hook := *r.Parent.Item.(*github.Hook)
	delivery := r.Item.(*github.HookDelivery)
	c := meta.(*client.Client)

	deliveryWithRequestResponse, _, err := c.Github.Organizations.GetHookDelivery(ctx, c.Org, *hook.ID, *delivery.ID)
	if err != nil {
		return err
	}

	r.SetItem(deliveryWithRequestResponse)
	return nil
}
