package promotion_codes

import (
	"context"

	"fmt"
	"strconv"

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
		Resolver:    fetchPromotionCodes("promotion_codes"),

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

func fetchPromotionCodes(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.PromotionCodeListParams{}

		if cl.Backend != nil {
			value, err := cl.Backend.Get(ctx, tableName, cl.ID())
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}
			if value != "" {
				vi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("retrieved invalid state backend: %q %w", value, err)
				}
				lp.Created = &vi
			}
		}

		it := cl.Services.PromotionCodes.List(lp)
		for it.Next() {
			res <- it.PromotionCode()
		}
		return it.Err()
	}
}
