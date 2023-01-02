package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
)

func OrderResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &shopify.Order{},
			Service:    "order",
			PKColumns:  []string{"id"},
		},
	}
}
