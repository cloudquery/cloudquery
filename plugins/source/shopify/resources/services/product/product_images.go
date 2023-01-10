package product

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ProductImages() *schema.Table {
	return &schema.Table{
		Name:      "shopify_product_images",
		Resolver:  fetchProductImages,
		Transform: transformers.TransformWithStruct(&shopify.ProductImage{}),
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
