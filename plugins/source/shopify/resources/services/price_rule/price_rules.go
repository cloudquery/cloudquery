package price_rule

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PriceRules() *schema.Table {
	return &schema.Table{
		Name:      "shopify_price_rules",
		Resolver:  fetchPriceRules,
		Transform: transformers.TransformWithStruct(&shopify.PriceRule{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:           "updated_at",
				Type:           arrow.FixedWidthTypes.Timestamp_us,
				Resolver:       schema.PathResolver("UpdatedAt"),
				IncrementalKey: true,
			},
		},
		IsIncremental: true,

		Relations: []*schema.Table{
			PriceRuleDiscountCodes(),
		},
	}
}
