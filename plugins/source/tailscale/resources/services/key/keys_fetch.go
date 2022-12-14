package key

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func fetchKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	result, err := c.Keys(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}

func getKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	key := resource.Item.(tailscale.Key)

	result, err := c.GetKey(ctx, key.ID)
	if err != nil {
		return err
	}

	resource.SetItem(result)
	return nil
}
