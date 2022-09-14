package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
)

func Container() []Resource {
	var registryRelations = []resourceDefinition{
		{
			azureStruct:  &containerregistry.Replication{},
			listFunction: "List",
			listFunctionArgsInit: []string{`registry := parent.Item.(containerregistry.Registry)
			resource, err := client.ParseResourceID(*registry.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*registry.Name"},
			isRelation:               true,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice",
						"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry",
					},
				},
			},
			definitions: append([]resourceDefinition{
				{
					azureStruct:  &containerservice.ManagedCluster{},
					listFunction: "List",
				},
				{
					azureStruct:  &containerregistry.Registry{},
					listFunction: "List",
					relations:    registryRelations,
				},
			}, registryRelations...),
			serviceNameOverride: "Container",
		},
	}

	return generateResources(resourcesByTemplates)
}
