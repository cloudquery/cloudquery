package billing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func ActionBillings() *schema.Table {
	return &schema.Table{
		Name:        "github_action_billing",
		Description: "ActionBilling represents a GitHub Action billing.",
		Resolver:    fetchActionBillings,
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
				Name: "total_minutes_used",
				Type: schema.TypeInt,
			},
			{
				Name: "total_paid_minutes_used",
				Type: schema.TypeFloat,
			},
			{
				Name: "included_minutes",
				Type: schema.TypeInt,
			},
			{
				Name: "minutes_used_breakdown",
				Type: schema.TypeJSON,
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
		return errors.WithStack(err)
	}
	res <- billing
	return nil
}
