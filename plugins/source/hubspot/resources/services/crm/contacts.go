package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Contacts() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_contacts",
		Description: "https://developers.hubspot.com/docs/api/crm/contacts",
		Transform: transformers.TransformWithStruct(
			contacts.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
		Resolver: fetchContacts,
	}
}
