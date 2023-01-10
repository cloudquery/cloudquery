package treasury

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryTransactionEntries() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_transaction_entries",
		Description: `https://stripe.com/docs/api/treasury_transaction_entries`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryTransactionEntry{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryTransactionEntries,

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

func fetchTreasuryTransactionEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	it := cl.Services.TreasuryTransactionEntries.List(&stripe.TreasuryTransactionEntryListParams{
		FinancialAccount: stripe.String(p.ID),
	})
	for it.Next() {
		res <- it.TreasuryTransactionEntry()
	}
	return it.Err()
}
