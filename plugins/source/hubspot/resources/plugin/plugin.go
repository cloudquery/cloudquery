package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/services/crm"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"crm":     "CRM",
	"hubspot": "HubSpot",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range source.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	return csr.ToTitle(table.Name)
}

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
		source.WithTitleTransformer(titleTransformer),
	)
}
