package tax

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TaxCodes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_tax_codes",
		Description: `https://stripe.com/docs/api/tax_codes`,
		Transform:   transformers.TransformWithStruct(&stripe.TaxCode{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTaxCodes,

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

func fetchTaxCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TaxCodeListParams{}

	it := cl.Services.TaxCodes.List(lp)
	for it.Next() {
		res <- it.TaxCode()
	}

	return it.Err()
}
