package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func CDN() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &cdn.Profile{},
					listFunction: "List",
				},
			},
			serviceNameOverride: "CDN",
		},
	}

	return generateResources(resourcesByTemplates)
}
