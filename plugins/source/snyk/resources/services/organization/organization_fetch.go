package organization

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, _, err := c.Orgs.List(ctx)
	if err != nil {
		return err
	}

	// limit orgs
	for _, org := range result {
		if c.WantOrganization(org.ID) {
			res <- org
		}
	}

	return nil
}
