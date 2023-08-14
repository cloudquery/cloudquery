package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/products"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_products",
		Description: "https://developers.hubspot.com/docs/api/crm/products",
		Resolver:    fetchProducts,
		Transform: transformers.TransformWithStruct(
			products.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
