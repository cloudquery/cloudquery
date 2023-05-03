package checkout

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/stripe/stripe-go/v74"
)

func CheckoutSessions() *schema.Table {
	return &schema.Table{
		Name:        "stripe_checkout_sessions",
		Description: `https://stripe.com/docs/api/checkout/sessions`,
		Transform:   client.TransformWithStruct(&stripe.CheckoutSession{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchCheckoutSessions,

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

		Relations: []*schema.Table{
			CheckoutSessionLineItems(),
		},
	}
}

func fetchCheckoutSessions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.CheckoutSessionListParams{}

	it := cl.Services.CheckoutSessions.List(lp)
	for it.Next() {
		res <- it.CheckoutSession()
	}

	return it.Err()
}
