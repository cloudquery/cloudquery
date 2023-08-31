package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/go-github/v49/github"
)

func Action() *schema.Table {
	return &schema.Table{
		Name:      "github_billing_action",
		Resolver:  fetchAction,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.ActionBilling{}),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchAction(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetActionsBillingOrg(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- billing
	return nil
}
