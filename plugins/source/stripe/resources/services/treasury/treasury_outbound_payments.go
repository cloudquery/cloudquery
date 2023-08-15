package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryOutboundPayments() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_outbound_payments",
		Description: `https://stripe.com/docs/api/treasury/outbound_payments`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryOutboundPayment{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryOutboundPayments,

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

func fetchTreasuryOutboundPayments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryOutboundPaymentListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryOutboundPayments.List(lp)
	for it.Next() {
		res <- it.TreasuryOutboundPayment()
	}

	return it.Err()
}
