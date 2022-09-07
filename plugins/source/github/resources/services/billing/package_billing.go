package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func PackageBillings() *schema.Table {

	return &schema.Table{
		Name:        "github_package_billing",
		Description: "PackageBilling represents a GitHub Package billing.",
		Resolver:    fetchPackageBillings,
		Multiplex:   client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:            "org",
				Description:     "The Github Organization of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveOrg,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "total_gigabytes_bandwidth_used",
				Type: schema.TypeInt,
			},
			{
				Name: "total_paid_gigabytes_bandwidth_used",
				Type: schema.TypeInt,
			},
			{
				Name: "included_gigabytes_bandwidth",
				Type: schema.TypeInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchPackageBillings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetPackagesBillingOrg(ctx, c.Org)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- billing
	return nil
}
