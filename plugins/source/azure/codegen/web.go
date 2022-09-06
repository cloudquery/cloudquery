package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Web() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &web.Site{},
					listFunction:       "List",
					subServiceOverride: "Apps",
					mockListResult:     "AppCollection",
					relations:          []string{"SiteAuthSettings(),PublishingProfiles()"},
					customColumns:      codegen.ColumnDefinitions{codegen.ColumnDefinition{Name: "vnet_connection", Type: schema.TypeJSON, Resolver: "fetchVnetConnections"}},
				},
				{
					azureStruct:          &web.SiteAuthSettings{},
					listFunction:         "GetAuthSettings",
					listFunctionArgsInit: []string{"p := parent.Item.(web.Site)"},
					listFunctionArgs:     []string{"*p.ResourceGroup", "*p.Name"},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation: true,
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
