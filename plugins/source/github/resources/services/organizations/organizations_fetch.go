package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	org, _, err := c.Github.Organizations.Get(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- org
	return nil
}
