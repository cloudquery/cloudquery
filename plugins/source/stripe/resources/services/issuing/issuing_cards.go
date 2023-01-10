package issuing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IssuingCards() *schema.Table {
	return &schema.Table{
		Name:        "stripe_issuing_cards",
		Description: `https://stripe.com/docs/api/issuing_cards`,
		Transform:   transformers.TransformWithStruct(&stripe.IssuingCard{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIssuingCards,

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

func fetchIssuingCards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IssuingCards.List(&stripe.IssuingCardListParams{})
	for it.Next() {
		res <- it.IssuingCard()
	}
	return it.Err()
}
