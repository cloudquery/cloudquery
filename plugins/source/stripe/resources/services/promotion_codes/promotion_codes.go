package promotion_codes

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PromotionCodes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_promotion_codes",
		Description: `https://stripe.com/docs/api/promotion_codes`,
		Transform:   transformers.TransformWithStruct(&stripe.PromotionCode{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPromotionCodes,

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

func fetchPromotionCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.PromotionCodes.List(&stripe.PromotionCodeListParams{})
	for it.Next() {
		res <- it.PromotionCode()
	}
	return it.Err()
}
