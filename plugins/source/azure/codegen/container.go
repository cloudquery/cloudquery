package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
)

func Container() []Resource {
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
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice",
						"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &containerservice.ManagedCluster{},
					listFunction: "List",
				},
				{
					azureStruct:  &containerregistry.Registry{},
					listFunction: "List",
				},
			},
			serviceNameOverride: "Container",
		},
	}

	return generateResources(resourcesByTemplates)
}
