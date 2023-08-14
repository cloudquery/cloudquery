package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/deals"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Deals() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_deals",
		Resolver:    fetchDeals,
		Description: "https://developers.hubspot.com/docs/api/crm/deals",
		Transform: transformers.TransformWithStruct(
			deals.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
