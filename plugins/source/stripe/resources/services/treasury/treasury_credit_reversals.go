package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryCreditReversals() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_credit_reversals",
		Description: `https://stripe.com/docs/api/treasury/credit_reversals`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryCreditReversal{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryCreditReversals,

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

func fetchTreasuryCreditReversals(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryCreditReversals.List(lp)
	for it.Next() {
		res <- it.TreasuryCreditReversal()
	}

	return it.Err()
}
