package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryDebitReversals() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_debit_reversals",
		Description: `https://stripe.com/docs/api/treasury/debit_reversals`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryDebitReversal{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryDebitReversals,

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

func fetchTreasuryDebitReversals(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryDebitReversalListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryDebitReversals.List(lp)
	for it.Next() {
		res <- it.TreasuryDebitReversal()
	}

	return it.Err()
}
