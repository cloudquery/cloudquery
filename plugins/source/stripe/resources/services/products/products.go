package products

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "stripe_products",
		Description: `https://stripe.com/docs/api/products`,
		Transform:   transformers.TransformWithStruct(&stripe.Product{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("Attributes", "DeactivateOn"))),
		Resolver:    fetchProducts,

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

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Products.List(&stripe.ProductListParams{})
	for it.Next() {
		res <- it.Product()
	}
	return it.Err()
}
