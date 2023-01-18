package product

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ProductVariants() *schema.Table {
	return &schema.Table{
		Name:      "shopify_product_variants",
		Resolver:  fetchProductVariants,
		Transform: transformers.TransformWithStruct(&shopify.ProductVariant{}),
		Columns: []schema.Column{
			{
				Name:     "product_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProductID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
