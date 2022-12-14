package device

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func fetchRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	device := parent.Item.(tailscale.Device)

	result, err := c.DeviceSubnetRoutes(ctx, device.ID)
	if err != nil {
		return err
	}

	res <- result
	return nil
}
