package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryTransactions() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_transactions",
		Description: `https://stripe.com/docs/api/treasury_transactions`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryTransaction{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryTransactions,

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

func fetchTreasuryTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	it := cl.Services.TreasuryTransactions.List(&stripe.TreasuryTransactionListParams{
		FinancialAccount: stripe.String(p.ID),
	})
	for it.Next() {
		res <- it.TreasuryTransaction()
	}
	return it.Err()
}
