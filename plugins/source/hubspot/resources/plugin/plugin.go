package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/services/crm"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cloudquery-hubspot",
		Version,
		schema.Tables{
			crm.Contacts(),
			crm.Companies(),
			crm.Deals(),
			crm.LineItems(),
			crm.Products(),
			crm.Tickets(),
			crm.Quotes(),
			crm.Owners(),
			crm.Pipelines(),
		},
		client.New,
	)
}
