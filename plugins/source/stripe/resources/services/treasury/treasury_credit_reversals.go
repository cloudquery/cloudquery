package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryCreditReversals() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_credit_reversals",
		Description: `https://stripe.com/docs/api/treasury_credit_reversals`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryCreditReversal{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTreasuryCreditReversals,

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
