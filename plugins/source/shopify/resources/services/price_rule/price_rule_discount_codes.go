package price_rule

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PriceRuleDiscountCodes() *schema.Table {
	return &schema.Table{
		Name:      "shopify_price_rule_discount_codes",
		Resolver:  fetchPriceRuleDiscountCodes,
		Transform: transformers.TransformWithStruct(&shopify.DiscountCode{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "price_rule_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PriceRuleID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
