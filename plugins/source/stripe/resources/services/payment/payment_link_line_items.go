package payment

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentLinkLineItems() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_link_line_items",
		Description: `https://stripe.com/docs/api/payment_links/line_items`,
		Transform:   client.TransformWithStruct(&stripe.LineItem{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPaymentLinkLineItems,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "payment_link_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchPaymentLinkLineItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.PaymentLink)

	lp := &stripe.PaymentLinkListLineItemsParams{
		PaymentLink: stripe.String(p.ID),
	}

	it := cl.Services.PaymentLinks.ListLineItems(lp)
	for it.Next() {
		res <- it.LineItem()
	}

	return it.Err()
}
