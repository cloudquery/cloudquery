package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
)

func CheckoutResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &shopify.Checkout{},
			Service:    "checkout",
			TableName:  "abandoned_checkouts",
			PKColumns:  []string{"id"},
		},
	}
}
