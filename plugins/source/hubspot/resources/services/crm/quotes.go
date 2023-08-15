package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/quotes"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Quotes() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_quotes",
		Resolver:    fetchQuotes,
		Description: "https://developers.hubspot.com/docs/api/crm/quotes",
		Transform: transformers.TransformWithStruct(
			quotes.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
