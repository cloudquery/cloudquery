package issuing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IssuingTransactions() *schema.Table {
	return &schema.Table{
		Name:        "stripe_issuing_transactions",
		Description: `https://stripe.com/docs/api/issuing_transactions`,
		Transform:   transformers.TransformWithStruct(&stripe.IssuingTransaction{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIssuingTransactions,

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

func fetchIssuingTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IssuingTransactions.List(&stripe.IssuingTransactionListParams{})
	for it.Next() {
		res <- it.IssuingTransaction()
	}
	return it.Err()
}
