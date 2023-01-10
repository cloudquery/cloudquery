package transfers

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Transfers() *schema.Table {
	return &schema.Table{
		Name:        "stripe_transfers",
		Description: `https://stripe.com/docs/api/transfers`,
		Transform:   transformers.TransformWithStruct(&stripe.Transfer{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTransfers,

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

func fetchTransfers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Transfers.List(&stripe.TransferListParams{})
	for it.Next() {
		res <- it.Transfer()
	}
	return it.Err()
}
