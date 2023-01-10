package price_rule

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PriceRules() *schema.Table {
	return &schema.Table{
		Name:      "shopify_price_rules",
		Resolver:  fetchPriceRules,
		Transform: transformers.TransformWithStruct(&shopify.PriceRule{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			PriceRuleDiscountCodes(),
		},
	}
}
