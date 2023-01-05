package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
)

func PriceRuleResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &shopify.PriceRule{},
			Service:    "price_rule",
			PKColumns:  []string{"id"},
			Relations:  []string{"PriceRuleDiscountCodes()"},
		},
		{
			DataStruct: &shopify.DiscountCode{},
			Service:    "price_rule",
			PKColumns:  []string{"price_rule_id", "id"},
		},
	}
}
