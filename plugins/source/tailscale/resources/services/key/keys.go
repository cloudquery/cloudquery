package key

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:                "tailscale_keys",
		Description:         `https://github.com/tailscale/tailscale/blob/main/api.md#keys`,
		Resolver:            fetchKeys,
		PreResourceResolver: getKey,
		Transform:           transformers.TransformWithStruct(&tailscale.Key{}, client.SharedTransformers(transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{
			{
				Name:     "tailnet",
				Type:     schema.TypeString,
				Resolver: client.ResolveTailnet,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.TailscaleClient.Keys(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}

func getKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	key := resource.Item.(tailscale.Key)
	result, err := c.TailscaleClient.GetKey(ctx, key.ID)
	if err != nil {
		return err
	}

	resource.SetItem(result)
	return nil
}
