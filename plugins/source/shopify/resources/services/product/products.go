package product

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:      "shopify_products",
		Resolver:  fetchProducts,
		Transform: transformers.TransformWithStruct(&shopify.Product{}, transformers.WithSkipFields("Variants", "Images")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
		IsIncremental: true,

		Relations: []*schema.Table{
			ProductVariants(),
			ProductImages(),
		},
	}
}
