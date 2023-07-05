package product

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ProductImages() *schema.Table {
	return &schema.Table{
		Name:      "shopify_product_images",
		Resolver:  fetchProductImages,
		Transform: transformers.TransformWithStruct(&shopify.ProductImage{}),
		Columns: []schema.Column{
			{
				Name:       "product_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ProductID"),
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
