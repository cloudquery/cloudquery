package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config billing.hcl --output .
func ActionBillings() *schema.Table {
	return &schema.Table{
		Name:        "github_action_billing",
		Description: "ActionBilling represents a GitHub Action billing.",
		Resolver:    fetchActionBillings,
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
				Name: "total_minutes_used",
				Type: schema.TypeBigInt,
			},
			{
				Name: "total_paid_minutes_used",
				Type: schema.TypeFloat,
			},
			{
				Name: "included_minutes",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "minutes_used_breakdown_ubuntu",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MinutesUsedBreakdown.Ubuntu"),
			},
			{
				Name:     "minutes_used_breakdown_mac_os",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MinutesUsedBreakdown.MacOS"),
			},
			{
				Name:     "minutes_used_breakdown_windows",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MinutesUsedBreakdown.Windows"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchActionBillings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetActionsBillingOrg(ctx, c.Org)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- billing
	return nil
}
