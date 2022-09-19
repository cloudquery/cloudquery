package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
)

func Web() []Resource {
	var authSettingsResource = resourceDefinition{
		azureStruct:          &web.SiteAuthSettings{},
		listFunction:         "GetAuthSettings",
		listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
		listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name"},
		listHandler: `if err != nil {
				return err
			}
			res <- response`,
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`},
		mockListResult:           mockDirectResponse,
	}

	var vnetInfoResource = resourceDefinition{
		azureStruct:          &web.VnetInfo{},
		listFunction:         "GetVnetConnection",
		listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
		listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name", "*site.SiteConfig.VnetName"},
		listHandler: `if err != nil {
				return err
			}
			res <- response`,
		subServiceOverride:       "VnetConnections",
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		mockListResult:           mockDirectResponse,
	}
	var publishingProfileResource = resourceDefinition{
		azureStruct:          &services.PublishingProfile{},
		listFunction:         `ListPublishingProfiles`,
		listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
		listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name"},
		listHandler: `if err != nil {
			return err
		}
	
		res <- response`,
		mockListFunctionArgsInit: []string{""},
		mockListFunctionArgs:     []string{`"test"`, `"test"`},
		mockListResult:           mockDirectResponse,
		mockDefinitionType:       "PublishingProfiles",
	}

	var appRelations = []resourceDefinition{
		authSettingsResource,
		vnetInfoResource,
		publishingProfileResource,
	}

	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &web.Site{},
					listFunction:       "List",
					subServiceOverride: "Apps",
					mockListResult:     "AppCollection",
					relations:          appRelations,
					mockListFunctionArgsInit: []string{
						`vnetName := "test"`,
						`result.Values()[0].SiteConfig.VnetName = &vnetName`,
						`resourceGroup := "test"`,
						`result.Values()[0].ResourceGroup = &resourceGroup`,
					},
				},
			},
			serviceNameOverride: "Web",
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, appRelations...)

	return generateResources(resourcesByTemplates)
}
