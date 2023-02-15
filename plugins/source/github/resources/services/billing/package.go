package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Package() *schema.Table {
	return &schema.Table{
		Name:      "github_billing_package",
		Resolver:  fetchPackage,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.PackageBilling{}, client.SharedTransformers()...),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchPackage(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetPackagesBillingOrg(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- billing
	return nil
}
