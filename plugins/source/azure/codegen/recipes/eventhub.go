package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
)

func EventHub() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &eventhub.EHNamespace{},
					listFunction:       "List",
					subServiceOverride: "Namespaces",
				},
			},
			serviceNameOverride: "EventHub",
		},
	}

	return generateResources(resourcesByTemplates)
}
