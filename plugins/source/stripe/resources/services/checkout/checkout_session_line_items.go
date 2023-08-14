package checkout

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func CheckoutSessionLineItems() *schema.Table {
	return &schema.Table{
		Name:        "stripe_checkout_session_line_items",
		Description: `https://stripe.com/docs/api/checkout/sessions/line_items`,
		Transform:   client.TransformWithStruct(&stripe.LineItem{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchCheckoutSessionLineItems,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:       "session_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchCheckoutSessionLineItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.CheckoutSession)

	lp := &stripe.CheckoutSessionListLineItemsParams{
		Session: stripe.String(p.ID),
	}

	it := cl.Services.CheckoutSessions.ListLineItems(lp)
	for it.Next() {
		res <- it.LineItem()
	}

	return it.Err()
}
