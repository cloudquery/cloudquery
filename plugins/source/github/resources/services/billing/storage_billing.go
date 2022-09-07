package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

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
				Type: schema.TypeInt,
			},
			{
				Name: "estimated_paid_storage_for_month",
				Type: schema.TypeFloat,
			},
			{
				Name: "estimated_storage_for_month",
				Type: schema.TypeInt,
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
		return errors.WithStack(err)
	}
	res <- billing
	return nil
}
