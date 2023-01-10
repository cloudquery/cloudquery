package prices

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Prices() *schema.Table {
	return &schema.Table{
		Name:        "stripe_prices",
		Description: `https://stripe.com/docs/api/prices`,
		Transform:   transformers.TransformWithStruct(&stripe.Price{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPrices,

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

func fetchPrices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Prices.List(&stripe.PriceListParams{})
	for it.Next() {
		res <- it.Price()
	}
	return it.Err()
}
