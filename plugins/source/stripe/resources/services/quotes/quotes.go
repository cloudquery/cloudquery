package quotes

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Quotes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_quotes",
		Description: `https://stripe.com/docs/api/quotes`,
		Transform:   client.TransformWithStruct(&stripe.Quote{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchQuotes,

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

func fetchQuotes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.QuoteListParams{}

	it := cl.Services.Quotes.List(lp)
	for it.Next() {
		res <- it.Quote()
	}

	return it.Err()
}
