package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryReceivedDebits() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_received_debits",
		Description: `https://stripe.com/docs/api/treasury_received_debits`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryReceivedDebit{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTreasuryReceivedDebits,

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

func fetchTreasuryReceivedDebits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryReceivedDebitListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryReceivedDebits.List(lp)
	for it.Next() {
		res <- it.TreasuryReceivedDebit()
	}

	return it.Err()
}
