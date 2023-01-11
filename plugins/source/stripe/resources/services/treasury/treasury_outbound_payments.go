package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryOutboundPayments() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_outbound_payments",
		Description: `https://stripe.com/docs/api/treasury_outbound_payments`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryOutboundPayment{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTreasuryOutboundPayments,

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
