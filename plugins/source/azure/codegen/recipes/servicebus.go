package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func ServiceBus() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &servicebus.SBNamespace{},
					listFunction:       "List",
					subServiceOverride: "Namespaces",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
