package reviews

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Reviews() *schema.Table {
	return &schema.Table{
		Name:        "stripe_reviews",
		Description: `https://stripe.com/docs/api/reviews`,
		Transform:   transformers.TransformWithStruct(&stripe.Review{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchReviews,

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

func fetchReviews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Reviews.List(&stripe.ReviewListParams{})
	for it.Next() {
		res <- it.Review()
	}
	return it.Err()
}
