package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
)

func IotHub() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &devices.IotHubDescription{},
					listFunction:       "ListBySubscription",
					subServiceOverride: "Devices",
				},
			},
			serviceNameOverride: "IotHub",
		},
	}

	return generateResources(resourcesByTemplates)
}
