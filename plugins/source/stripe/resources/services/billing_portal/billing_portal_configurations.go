package billing_portal

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func BillingPortalConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_billing_portal_configurations",
		Description: `https://stripe.com/docs/api/customer_portal/configuration`,
		Transform:   client.TransformWithStruct(&stripe.BillingPortalConfiguration{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchBillingPortalConfigurations,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
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
