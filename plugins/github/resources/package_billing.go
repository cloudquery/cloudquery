package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config billing.hcl --output .
func PackageBillings() *schema.Table {
	return &schema.Table{
		Name:        "github_package_billing",
		Description: "PackageBilling represents a GitHub Package billing.",
		Resolver:    fetchPackageBillings,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"org"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
			},
			{
				Name: "total_gigabytes_bandwidth_used",
				Type: schema.TypeBigInt,
			},
			{
				Name: "total_paid_gigabytes_bandwidth_used",
				Type: schema.TypeBigInt,
			},
			{
				Name: "included_gigabytes_bandwidth",
				Type: schema.TypeBigInt,
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
		return diag.WrapError(err)
	}
	res <- billing
	return nil
}
