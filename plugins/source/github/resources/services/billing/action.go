package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Action() *schema.Table {
	return &schema.Table{
		Name:      "github_billing_action",
		Resolver:  fetchAction,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.ActionBilling{}, client.SharedTransformers()...),
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
