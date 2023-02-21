package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryDebitReversals() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_debit_reversals",
		Description: `https://stripe.com/docs/api/treasury_debit_reversals`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryDebitReversal{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTreasuryDebitReversals,

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
