package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func Storage() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &storage.Account{},
					listFunction: "List",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
