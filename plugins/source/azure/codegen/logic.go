package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
)

func Logic() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &logic.Workflow{},
					listFunction:         "ListBySubscription",
					listFunctionArgs:     []string{"&top", "\"\""},
					listFunctionArgsInit: []string{"var top int32 = 100"},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
