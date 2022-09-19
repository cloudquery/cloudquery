package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
)

func Search() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &search.Service{},
					listFunction:     "ListBySubscription",
					listFunctionArgs: []string{"nil"},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
