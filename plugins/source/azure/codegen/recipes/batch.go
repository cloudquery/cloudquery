package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
)

func Batch() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &batch.Account{},
					listFunction: "List",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
