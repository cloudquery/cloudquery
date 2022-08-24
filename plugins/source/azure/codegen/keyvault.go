package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
)

func KeyValueResources() []Resource {
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
					azureStruct:          &keyvault.ManagedHsm{},
					listFunction:         "ListBySubscription",
					listFunctionArgs:     []string{"&maxResults"},
					listFunctionArgsInit: []string{"maxResults := int32(100)"},
					subServiceOverride:   "ManagedHSMs",
				},
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
