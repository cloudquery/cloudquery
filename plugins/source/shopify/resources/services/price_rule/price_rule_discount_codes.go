package price_rule

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PriceRuleDiscountCodes() *schema.Table {
	return &schema.Table{
		Name:      "shopify_price_rule_discount_codes",
		Resolver:  fetchPriceRuleDiscountCodes,
		Transform: transformers.TransformWithStruct(&shopify.DiscountCode{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:       "price_rule_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("PriceRuleID"),
				PrimaryKey: true,
			},
		},
	}
}
