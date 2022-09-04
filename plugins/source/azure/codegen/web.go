package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

const fetchVnetConnections = `func fetchVnetConnections(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	site := resource.Item.(web.Site)
	svc := meta.(*client.Client).Services().Web.Apps

	if site.SiteConfig == nil || site.SiteConfig.VnetName == nil {
		return nil
	}

	response, err := svc.GetVnetConnection(ctx, *site.ResourceGroup, *site.Name, *site.SiteConfig.VnetName)
	if err != nil {
		return diag.WrapError(err)
	}
	if response.VnetInfoProperties != nil {
		vnetConnection := make(map[string]interface{})
		if response.Name != nil {
			vnetConnection["name"] = response.Name
		}
		if response.ID != nil {
			vnetConnection["id"] = response.ID
		}
		if response.Type != nil {
			vnetConnection["type"] = response.Type
		}
		vnetConnection["properties"] = response.VnetInfoProperties
		b, err := json.Marshal(vnetConnection)
		if err != nil {
			return diag.WrapError(err)
		}
		return diag.WrapError(resource.Set(c.Name, b))
	}
	return nil
}`

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
					relations:          []string{"SiteAuthSettings()"},
					helpers:            []string{fetchVnetConnections},
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
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
