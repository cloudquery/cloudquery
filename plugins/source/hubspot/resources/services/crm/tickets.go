package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/tickets"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Tickets() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_tickets",
		Resolver:    fetchTickets,
		Description: "https://developers.hubspot.com/docs/api/crm/tickets",
		Transform: transformers.TransformWithStruct(
			tickets.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
