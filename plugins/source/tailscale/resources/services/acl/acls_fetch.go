package acl

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAcls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.ACL(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}
