package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
)

func StreamAnalytics() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &streamanalytics.StreamingJob{},
					listFunction:     "List",
					listFunctionArgs: []string{`""`},
					skipFields:       []string{"Properties", "Datasource", "Serialization"},
				},
			},
			serviceNameOverride: "StreamAnalytics",
		},
	}

	return generateResources(resourcesByTemplates)
}
