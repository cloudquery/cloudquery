package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	hsmKeyValue "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
)

func KeyValue() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &hsmKeyValue.ManagedHsm{},
					listFunction:         "ListBySubscription",
					listFunctionArgs:     []string{"&maxResults"},
					listFunctionArgsInit: []string{"maxResults := int32(100)"},
				},
			},
			serviceNameOverride: "KeyVault",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &keyvault.Vault{},
					listFunction:         "ListBySubscription",
					listFunctionArgs:     []string{"&maxResults"},
					listFunctionArgsInit: []string{"maxResults := int32(1000)"},
				},
			},
			serviceNameOverride: "KeyVault",
		},
	}

	return generateResources(resourcesByTemplates)
}
