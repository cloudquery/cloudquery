package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/line_items"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LineItems() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_line_items",
		Description: "https://developers.hubspot.com/docs/api/crm/line-items",
		Resolver:    fetchLineItems,
		Transform: transformers.TransformWithStruct(
			line_items.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
