package recipes

import (
	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
)

func Logic() []Resource {
	var workflowRelations = []resourceDefinition{
		{
			azureStruct:              &insights.DiagnosticSettingsResource{},
			listFunction:             "List",
			listHandler:              valueHandler,
			listFunctionArgs:         []string{"*workflow.ID"},
			listFunctionArgsInit:     []string{"workflow := parent.Item.(logic.Workflow)"},
			subServiceOverride:       "DiagnosticSettings",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"/subscriptions/test/resourceGroups/test/providers/test/test/test"`},
			mockListResult:           "DiagnosticSettingsResourceCollection",
		},
	}
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
					relations:            workflowRelations,
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"},
				},
			},
			definitions:         workflowRelations,
			serviceNameOverride: "Logic",
		},
	}

	initParents(resourcesByTemplates)
	return generateResources(resourcesByTemplates)
}
