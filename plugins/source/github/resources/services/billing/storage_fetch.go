package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchBillingStorage(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetStorageBillingOrg(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- billing
	return nil
}
