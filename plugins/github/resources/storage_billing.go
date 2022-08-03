package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config billing.hcl --output .
func StorageBillings() *schema.Table {
	return &schema.Table{
		Name:        "github_storage_billing",
		Description: "StorageBilling represents a GitHub Storage billing.",
		Resolver:    fetchStorageBillings,
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
				Name: "days_left_in_billing_cycle",
				Type: schema.TypeBigInt,
			},
			{
				Name: "estimated_paid_storage_for_month",
				Type: schema.TypeFloat,
			},
			{
				Name: "estimated_storage_for_month",
				Type: schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchStorageBillings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetStorageBillingOrg(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- billing
	return nil
}
