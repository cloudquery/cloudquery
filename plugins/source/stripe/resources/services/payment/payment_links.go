package payment

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentLinks() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_links",
		Description: `https://stripe.com/docs/api/payment_links`,
		Transform:   client.TransformWithStruct(&stripe.PaymentLink{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPaymentLinks,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			PaymentLinkLineItems(),
		},
	}
}

func fetchPaymentLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.PaymentLinkListParams{}

	it := cl.Services.PaymentLinks.List(lp)
	for it.Next() {
		res <- it.PaymentLink()
	}

	return it.Err()
}
