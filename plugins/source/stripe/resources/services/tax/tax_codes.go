package tax

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TaxCodes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_tax_codes",
		Description: `https://stripe.com/docs/api/tax_codes`,
		Transform:   client.TransformWithStruct(&stripe.TaxCode{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTaxCodes,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
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
