package codegen

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

func Subscriptions() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_pager.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_pager_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &armsubscriptions.Subscription{},
					mockListResult: " ",
				},
				{
					azureStruct:        &armsubscriptions.TenantIDDescription{},
					subServiceOverride: "Tenants",
				},
			},
			serviceNameOverride: "Subscriptions",
		},
	}

	return generateResources(resourcesByTemplates)
}
