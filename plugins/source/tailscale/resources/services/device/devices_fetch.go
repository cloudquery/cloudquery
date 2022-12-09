package device

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDevices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	result, err := c.Devices(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}
