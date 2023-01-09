package key

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:                "tailscale_keys",
		Description:         `https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#Key`,
		Resolver:            fetchKeys,
		PreResourceResolver: getKey,
		Transform:           transformers.TransformWithStruct(&tailscale.Key{}),
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
