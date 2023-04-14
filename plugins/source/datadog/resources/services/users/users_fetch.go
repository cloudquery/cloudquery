package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.UsersAPI.ListUsers(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
