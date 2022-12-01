package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
)

func FrontDoor() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &frontdoor.FrontDoor{},
					listFunction:       "List",
					subServiceOverride: "Doors",
					mockListResult:     "ListResult",
					mockFieldsToIgnore: []string{"RouteConfiguration"},
				},
			},
			serviceNameOverride: "FrontDoor",
		},
	}

	return generateResources(resourcesByTemplates)
}
