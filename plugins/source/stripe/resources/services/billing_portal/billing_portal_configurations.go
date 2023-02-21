package billing_portal

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func BillingPortalConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_billing_portal_configurations",
		Description: `https://stripe.com/docs/api/billing_portal_configurations`,
		Transform:   transformers.TransformWithStruct(&stripe.BillingPortalConfiguration{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchBillingPortalConfigurations,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchBillingPortalConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.BillingPortalConfigurationListParams{}

	it := cl.Services.BillingPortalConfigurations.List(lp)
	for it.Next() {
		res <- it.BillingPortalConfiguration()
	}

	return it.Err()
}
