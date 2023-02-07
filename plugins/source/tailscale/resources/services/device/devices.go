package device

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_devices",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-devices-get`,
		Resolver:    fetchDevices,
		Transform:   transformers.TransformWithStruct(&tailscale.Device{}, client.SharedTransformers(transformers.WithPrimaryKeys("ID"))...),
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
		Relations: []*schema.Table{
			routes(),
		},
	}
}
