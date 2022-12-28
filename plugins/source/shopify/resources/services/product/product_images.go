// Code generated by codegen; DO NOT EDIT.

package product

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProductImages() *schema.Table {
	return &schema.Table{
		Name:     "shopify_product_images",
		Resolver: fetchProductImages,
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
			{
				Name:     "position",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Position"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "alt",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Alt"),
			},
			{
				Name:     "width",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Width"),
			},
			{
				Name:     "height",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Height"),
			},
			{
				Name:     "src",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Src"),
			},
			{
				Name:     "variant_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VariantIDs"),
			},
			{
				Name:     "admin_graphql_api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdminGraphqlAPIID"),
			},
		},
	}
}
