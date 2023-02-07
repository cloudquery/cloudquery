package device

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:                 "tailscale_devices",
		Description:          `https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-devices-get`,
		Resolver:             fetchDevices,
		PostResourceResolver: postDeviceFetch,
		Transform:            transformers.TransformWithStruct(&tailscale.Device{}, client.SharedTransformers(transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{
			{
				Name:     "tailnet",
				Type:     schema.TypeString,
				Resolver: client.ResolveTailnet,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "advertised_routes",
				Type: schema.TypeStringArray,
			},
			{
				Name: "enabled_routes",
				Type: schema.TypeStringArray,
			},
		},
	}
}

func fetchDevices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.TailscaleClient.Devices(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}

func postDeviceFetch(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)

	device := resource.Item.(tailscale.Device)

	result, err := c.TailscaleClient.DeviceSubnetRoutes(ctx, device.ID)
	if err != nil {
		return err
	}
	if err := resource.Set("advertised_routes", result.Advertised); err != nil {
		return err
	}
	return resource.Set("enabled_routes", result.Enabled)
}
