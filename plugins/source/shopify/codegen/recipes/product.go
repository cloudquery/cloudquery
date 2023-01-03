package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
)

func ProductResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &shopify.Product{},
			Service:    "product",
			PKColumns:  []string{"id"},
			SkipFields: []string{"Variants", "Images"},
			Relations:  []string{"ProductVariants()", "ProductImages()"},
		},
		{
			DataStruct: &shopify.ProductVariant{},
			Service:    "product",
			PKColumns:  []string{"product_id", "id"},
		},
		{
			DataStruct: &shopify.ProductImage{},
			Service:    "product",
			PKColumns:  []string{"product_id", "id"},
		},
	}
}
