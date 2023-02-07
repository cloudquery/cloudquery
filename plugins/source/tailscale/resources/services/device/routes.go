package device

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func routes() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_device_routes",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#device-routes-get`,
		Resolver:    fetchRoutes,
		Transform:   transformers.TransformWithStruct(&tailscale.DeviceRoutes{}),
		Columns: []schema.Column{
			{
				Name:     "tailnet",
				Type:     schema.TypeString,
				Resolver: client.ResolveTailnet,
			},
			{
				Name:     "device_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
