package topups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Topups() *schema.Table {
	return &schema.Table{
		Name:        "stripe_topups",
		Description: `https://stripe.com/docs/api/topups`,
		Transform:   transformers.TransformWithStruct(&stripe.Topup{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("Source"))),
		Resolver:    fetchTopups,

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

func fetchTopups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Topups.List(&stripe.TopupListParams{})
	for it.Next() {
		res <- it.Topup()
	}
	return it.Err()
}
