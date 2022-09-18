package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
)

func Redis() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &redis.ResourceType{},
					listFunction:       "ListBySubscription",
					mockListResult:     "ListResult",
					subServiceOverride: "Caches",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
