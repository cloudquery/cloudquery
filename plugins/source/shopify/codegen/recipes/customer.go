package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
)

func CustomerResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &shopify.Customer{},
			Service:    "customer",
			PKColumns:  []string{"id"},
		},
	}
}
